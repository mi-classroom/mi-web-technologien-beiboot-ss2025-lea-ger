<script lang="ts">
  import type {FileItem} from "@/utils.js";
  import { onMount } from 'svelte';
  import { getAllTags, isWritable, getTagType, getTagOptions, type IPTCTag } from '../iptcTagHelpers';
  import MetadataField from './MetadataField.svelte';

  export let selected: FileItem;

  const apiUrl = import.meta.env.VITE_APP_API_URL;
  let fileMetadata: Record<string, string> | null = null;
  let originalMetadata: Record<string, string> | null = null;
  let editableMetadata: Record<string, string> = {};
  let iptcTags: IPTCTag[] = [];
  let hasChanges = false;
  let loading = true;

  async function fetchMetadata(): Promise<void> {
    try {
      const response = await fetch(`${apiUrl}/api/metadata/${selected.id}`);
      if (response.ok) {
        fileMetadata = await response.json();
        originalMetadata = { ...fileMetadata };
        editableMetadata = { ...fileMetadata };
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

  onMount(async () => {
    await fetchMetadata();
    await loadTags();
  });
</script>

<div class="border border-primary p-4 bg-base-100 rounded-sm">
    <button class="btn btn-primary mb-4" disabled={!hasChanges}>Speichern</button>
    <div class="mb-4">
        <img src={`${apiUrl}/assets/${selected?.id}`}
             alt="Preview"
             class="w-full object-contain rounded border"
        />
    </div>

    {#if loading}
      <p class="text-gray-500">Lade Metadaten...</p>
    {:else if fileMetadata}
        <div class="text-sm divide-y">
            {#each Object.entries(editableMetadata) as [key, value]}
                {#if iptcTags.length}
                  {#await Promise.resolve(iptcTags.find(t => t.name === key)) then tag}
                    <MetadataField
                      keyName={key}
                      {value}
                      originalValue={originalMetadata ? originalMetadata[key] : ''}
                      {tag}
                      onEdit={onEdit}
                    />
                  {/await}
                {/if}
            {/each}
        </div>
    {:else}
        <p class="text-gray-500">Lade Metadaten...</p>
    {/if}
</div>