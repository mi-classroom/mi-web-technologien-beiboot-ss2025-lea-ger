<script lang="ts">
  import type {FileItem, IPTCTag} from "@/utils.js";
  import {getAllTags, mostRelevantIPTCTags} from '@/utils.js';
  import {onMount} from 'svelte';
  import MetadataField from './MetadataField.svelte';
  import TagSelector from '@/components/TagSelector.svelte';
  import Toast from '@/components/Toast.svelte';

  export let selected: FileItem;

  const apiUrl = import.meta.env.VITE_APP_API_URL;
  let fileMetadata: Record<string, string> | null = null;
  let originalMetadata: Record<string, string> | null = null;
  let editableMetadata: Record<string, string> = {};
  let iptcTags: IPTCTag[] = [];
  let hasChanges = false;
  let loading = true;
  let collapseOpen = false;
  let saving = false;
  let toastOpen = false;
  let toastMsg = '';
  let toastType: 'success' | 'error' | 'info' | 'warning' = 'info';

  async function fetchMetadata(): Promise<void> {
    try {
      const response = await fetch(`${apiUrl}/api/metadata/${selected.id}`);
      if (response.ok) {
        fileMetadata = await response.json();
        originalMetadata = {...fileMetadata};
        editableMetadata = {...fileMetadata};
      } else {
        console.error("Fehler beim Abrufen der Metadaten");
      }
    } catch (error) {
      console.error("Fehler beim Abrufen der Metadaten", error);
    }
  }

  async function loadTags() {
    iptcTags = await getAllTags();
    loading = false;
  }

  function onEdit(key: string, value: string) {
    editableMetadata[key] = value;
    hasChanges = Object.keys(editableMetadata).some(k => editableMetadata[k] !== originalMetadata?.[k]);
  }

  function onRemove(key: string) {
    const { [key]: _, ...rest } = editableMetadata;
    editableMetadata = rest;
    hasChanges = true;
  }

  function getTag(tagName: string): IPTCTag | undefined {
    return iptcTags.find(t => t.name === tagName);
  }

  async function addTag(tagName: string) {
    if (!editableMetadata[tagName]) {
      editableMetadata = { ...editableMetadata, [tagName]: '' };
      hasChanges = true;
    }
  }

  async function saveTags(): Promise<void> {
    saving = true;
    try {
      const response = await fetch(`${apiUrl}/api/metadata/${selected.id}`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(editableMetadata),
      });
      if (response.ok) {
        originalMetadata = { ...editableMetadata };
        hasChanges = false;
        toastType = 'success';
        toastMsg = 'Metadaten erfolgreich gespeichert';
        toastOpen = true;
      } else {
        toastType = 'error';
        toastMsg = 'Speichern der Metadaten fehlgeschlagen';
        toastOpen = true;
      }
    } catch (error) {
      toastType = 'error';
      toastMsg = 'Speichern der Metadaten fehlgeschlagen';
      toastOpen = true;
    } finally {
      saving = false;
    }
  }

  $: relevantTagNames = mostRelevantIPTCTags.map(t => t.name);
  $: relevantTotal = mostRelevantIPTCTags.length;
  $: relevantFilled = relevantTagNames.filter(name => !!editableMetadata[name] && editableMetadata[name].toString().trim() !== '').length;
  $: relevantPercent = Math.round((relevantFilled / relevantTotal) * 100);

  onMount(async () => {
    await fetchMetadata();
    await loadTags();
  });
</script>

<div class="border border-primary p-4 bg-base-100 rounded-sm">
    <div class="flex z-10 justify-end items-center sticky top-0 bg-base-100 py-4">
        <button
                class="btn btn-primary"
                disabled={!hasChanges}
                class:loading={saving}
                on:click={saveTags}
        >
            Speichern
        </button>
    </div>

    <div>
        <img src={`${apiUrl}/assets/${selected?.id}`}
             alt="Preview"
             class="w-full object-contain rounded border"
        />
    </div>

    {#if loading}
        <p class="text-gray-500">Lade Metadaten...</p>
    {:else if fileMetadata}

        <div class="text-sm divide-y">
            <div class="py-4">
                <label class="block mb-1 text-sm font-semibold text-gray-700">
                    Wichtige Tags befüllt: {relevantFilled} / {relevantTotal}
                </label>
                <progress class="progress progress-primary w-full" value={relevantFilled} max={relevantTotal}></progress>
                <div class="text-xs text-right text-gray-500">{relevantPercent}%</div>
            </div>

            <TagSelector
                    label="Tag hinzufügen"
                    excludedTags={Object.keys(editableMetadata)}
                    onSelect={addTag}
            />

            {#each Object.entries(editableMetadata) as [key, value] (key)}
                {#if iptcTags.some(tag => tag.name === key)}
                    <MetadataField
                            keyName={key}
                            value={value}
                            originalValue={originalMetadata ? originalMetadata[key] : ''}
                            tag={getTag(key)}
                            onEdit={onEdit}
                            onRemove={onRemove}
                    />
                {/if}
            {/each}

            <div class="collapse collapse-arrow border border-base-300 bg-base-100 rounded-box mt-4">
                <input type="checkbox" bind:checked={collapseOpen}/>
                <div class="collapse-title">
                    Nicht-IPTC-Tags
                </div>
                <div class="collapse-content">
                    {#each Object.entries(editableMetadata) as [key, value] (key)}
                        {#if !iptcTags.some(tag => tag.name === key)}
                            <MetadataField
                                    keyName={key}
                                    value={value}
                                    originalValue={originalMetadata ? originalMetadata[key] : ''}
                                    tag={getTag(key)}
                                    onEdit={onEdit}
                            />
                        {/if}
                    {/each}
                </div>
            </div>
        </div>
    {:else}
        <p class="text-gray-500">Keine Metadaten gefunden.</p>
    {/if}
</div>

<Toast message={toastMsg} type={toastType} bind:open={toastOpen} duration={3000} />