<script lang="ts">
  import {type FileItem, getTagByName, type IPTCTag} from "@/utils.js";
  import {slide} from 'svelte/transition';
  import TagSelector from "@/components/TagSelector.svelte";
  import MetadataField from "@/components/MetadataField.svelte";
  import Toast from "@/components/Toast.svelte";
  import { saveMacro, type MergeStategy } from "@/macros.ts";

  export let onClose: () => void;
  export let onUpdated: () => void;
  export let selected: FileItem[];

  type TagEntry = { value: any, tag: IPTCTag, strategy: MergeStategy };
  let tags: TagEntry[] = [];
  let saving = false;
  let toastOpen = false;
  let toastMsg = '';
  let toastType: 'success' | 'error' | 'info' | 'warning' = 'info';
  let progressCurrent = 0;
  let progressTotal = 0;

  async function saveChanges() {
    // Note: Merge-Strategien werden primär für Makros verwendet.
    // Der direkte Speichern-Button überschreibt immer die Werte (aktuelles Verhalten).
    saving = true;
    progressCurrent = 0;
    progressTotal = selected.length;
    let errorCount = 0;

    for (const item of selected) {
      try {
        const editableMetadata: Record<string, any> = {};
        for (const t of tags) {
          editableMetadata[t.tag.name] = t.value;
        }
        const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/metadata/${item.id}`, {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(editableMetadata),
        });
        if (!response.ok) errorCount++;
      } catch (error) {
        errorCount++;
      } finally {
        progressCurrent++;
      }
    }

    saving = false;
    if (errorCount === 0) {
      toastType = 'success';
      toastMsg = 'Alle Metadaten erfolgreich gespeichert';
    } else if (errorCount < selected.length) {
      toastType = 'warning';
      toastMsg = `${selected.length - errorCount} von ${selected.length} Metadaten gespeichert, ${errorCount} Fehler.`;
    } else {
      toastType = 'error';
      toastMsg = 'Speichern der Metadaten fehlgeschlagen';
    }
    toastOpen = true;
    onUpdated();
    closeEditor();
  }

  async function addTag(tag: string) {
    const existingTag = Object.values(tags).find(t => t.tag.name === tag);
    if (!existingTag) {
      const tagObj = await getTagByName(tag);
      if (!tagObj) {
        console.error(`Tag "${tag}" nicht gefunden.`);
        return;
      }
      tags = [...tags, { value: '', tag: tagObj, strategy: 'overwrite' }];
    } else {
      console.warn(`Tag "${tag}" ist bereits ausgewählt.`);
    }
  }

  function editTag(tag: string, value: any) {
    tags = tags.map(t =>
      t.tag.name === tag ? { ...t, value } : t
    );
  }

  function setStrategy(tagName: string, strategy: MergeStategy) {
    tags = tags.map(t => t.tag.name === tagName ? { ...t, strategy } : t);
  }

  function handleStrategyChange(tagName: string) {
    return (e: Event) => {
      const val = (e.target as HTMLSelectElement).value as MergeStategy;
      setStrategy(tagName, val);
    };
  }

  function closeEditor() {
    onClose();
  }

  function saveAsMacro() {
    if (tags.length === 0) return;
    const name = window.prompt('Makro-Name eingeben:');
    if (!name) return;
    saveMacro({ name, changes: tags.map(t => ({ tag: t.tag.name, value: t.value, strategy: t.strategy }))});
    toastType = 'success';
    toastMsg = 'Makro gespeichert';
    toastOpen = true;
  }
</script>

<div
        class="fixed inset-0 bg-black/50 opacity-100 z-50  flex items-end"
>
    <div class="w-full bg-base-100 min-h-1/3 rounded-t-lg p-6"
         transition:slide|global={{ duration: 300 }}
    >
        <div class="container mx-auto">
            <div class="flex flex-wrap items-center justify-between gap-1">
                <h3 class="font-bold text-lg">Metadaten bearbeiten</h3>
                <div class="flex items-center justify-end gap-4 mt-6">
                    <button class="btn btn-secondary" on:click={saveAsMacro} disabled={tags.length === 0}>Als Makro speichern</button>
                    <button class="btn btn-primary"
                            disabled={tags.length === 0 || saving}
                            class:loading={saving}
                            on:click={saveChanges}
                    >Speichern</button>
                    <button class="btn" on:click={closeEditor}>Abbrechen</button>
                </div>
            </div>
            {#if saving}
              <div class="my-4">
                <progress class="progress progress-primary w-full" value={progressCurrent} max={progressTotal}></progress>
                <div class="text-xs text-right text-gray-500">{progressCurrent} / {progressTotal} gespeichert</div>
              </div>
            {/if}
            <div class="space-y-2 my-4">
                {#each tags as t ( t.tag.name )}
                  <div class="flex items-center gap-3">
                    <div class="flex-1">
                      <MetadataField
                        keyName={t.tag.name}
                        value={t.value}
                        originalValue={undefined}
                        tag={t.tag}
                        onEdit={editTag}
                        onRemove={(k) => {
                          tags = tags.filter(x => x.tag.name !== k);
                        }}
                      />
                    </div>
                    <div class="text-sm">
                      <label class="block text-xs mb-1">Merging-Strategie</label>
                      <select class="select select-bordered w-full" value={t.strategy} on:change={handleStrategyChange(t.tag.name)}>
                        <option value="overwrite">Immer überschreiben</option>
                        <option value="skipIfFilled">Überspringen, wenn befüllt</option>
                      </select>
                      <div class="text-sm opacity-60 mt-1">Nutze $$, um den aktuellen Feldwert einzufügen.</div>
                    </div>
                  </div>
                {/each}
            </div>

            <div class="flex items-center gap-3">
              <TagSelector
                      label="Tag hinzufügen"
                      onSelect={addTag}
              />
            </div>
        </div>
    </div>
</div>

<Toast message={toastMsg} type={toastType} bind:open={toastOpen} duration={3000} />