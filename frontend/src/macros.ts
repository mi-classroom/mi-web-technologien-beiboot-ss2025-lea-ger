// Macro utilities for IPTC metadata editor
// Stores and retrieves macros from localStorage and applies them to items

import type {FileItem} from "@/utils.js";

export type MergeStategy = "overwrite" | "skipIfFilled";

export type MacroChange = {
  tag: string;
  value: any;
  strategy?: MergeStategy;
};

export type Macro = {
  id: string;
  name: string;
  changes: MacroChange[];
  createdAt?: string;
};

const STORAGE_KEY = "iptc_macros";

function loadAll(): Macro[] {
  try {
    const raw = localStorage.getItem(STORAGE_KEY);
    if (!raw) return [];
    const arr = JSON.parse(raw);
    if (Array.isArray(arr)) return arr as Macro[];
  } catch (e) {
    console.warn("Failed to parse macros from localStorage", e);
  }
  return [];
}

function persistAll(list: Macro[]) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(list));
}

function uuid(): string {
  // Simple unique ID generator: timestamp + random
  return (
    Date.now().toString(36) +
    "-" +
    Math.random().toString(36).substr(2, 9)
  );
}


export function listMacros(): Macro[] {
  return loadAll();
}

export function getMacro(id: string): Macro | undefined {
  return loadAll().find(m => m.id === id);
}

export function saveMacro(macro: Omit<Macro, "id" | "createdAt"> & Partial<Pick<Macro, "id" | "createdAt">>): Macro {
  const list = loadAll();
  const now = new Date().toISOString();
  const id = macro.id || uuid();
  const toSave: Macro = {
    id,
    name: macro.name,
    changes: macro.changes.map(c => ({...c, strategy: c.strategy || "overwrite"})),
    createdAt: macro.createdAt || now
  };
  const idx = list.findIndex(m => m.id === id);
  if (idx >= 0) list[idx] = toSave; else list.push(toSave);
  persistAll(list);
  return toSave;
}

export function deleteMacro(id: string) {
  const list = loadAll().filter(m => m.id !== id);
  persistAll(list);
}

// Helper to resolve macro value placeholders. Supports $$ for current value.
function resolveValue(input: any, current: any): any {
  if (typeof input === "string") {
    // If input contains $$, replace with current value stringified
    if (input.includes("$$")) {
      const cur = current == null ? "" : (Array.isArray(current) ? current.join(",") : String(current));
      return input.replaceAll("$$", cur);
    }
    return input;
  }
  if (Array.isArray(input)) {
    return input.map(v => resolveValue(v, current));
  }
  // numbers/objects are passed through
  return input;
}

async function fetchMetadata(itemId: number): Promise<Record<string, any> | null> {
  try {
    const res = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/metadata/${itemId}`);
    if (!res.ok) return null;
    return await res.json();
  } catch (e) {
    console.warn("Failed to fetch metadata for", itemId, e);
    return null;
  }
}

export async function applyMacroToItem(item: FileItem, macro: Macro): Promise<boolean> {
  const current = await fetchMetadata(item.id);
  const editable: Record<string, any> = {};
  for (const ch of macro.changes) {
    const strategy: MergeStategy = ch.strategy || "overwrite";
    const existing = current ? current[ch.tag] : undefined;
    // skipIfFilled: if existing has a truthy/non-empty value, skip this field
    if (strategy === "skipIfFilled") {
      const hasValue = Array.isArray(existing) ? existing.length > 0 : existing != null && existing !== "";
      if (hasValue) continue;
    }
    editable[ch.tag] = resolveValue(ch.value, existing);
  }
  if (Object.keys(editable).length === 0) return true; // nothing to update
  try {
    const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/metadata/${item.id}`, {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(editable),
    });
    return response.ok;
  } catch (e) {
    return false;
  }
}

export async function applyMacroToItems(items: FileItem[], macro: Macro, onProgress?: (done: number, total: number) => void): Promise<{
  success: number,
  failed: number
}> {
  let success = 0, failed = 0, done = 0;
  for (const item of items) {
    const ok = await applyMacroToItem(item, macro);
    if (ok) success++; else failed++;
    done++;
    onProgress?.(done, items.length);
  }
  return {success, failed};
}
