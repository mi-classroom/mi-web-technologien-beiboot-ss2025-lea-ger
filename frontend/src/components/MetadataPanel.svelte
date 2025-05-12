<script lang="ts">
  import { onMount } from "svelte";
  import type { FileItem } from "@/utils.js";

  export let selected: FileItem;

  const apiUrl = import.meta.env.VITE_APP_API_URL;
  let fileMetadata: Record<string, string> | null = null;

  async function fetchMetadata(): Promise<void> {
    try {
      const response = await fetch(`${apiUrl}/api/metadata/${selected.id}`);
      if (response.ok) {
        fileMetadata = await response.json();
      } else {
        console.error("Fehler beim Abrufen der Metadaten");
      }
    } catch (error) {
      console.error("Fehler beim Abrufen der Metadaten", error);
    }
  }

  fetchMetadata();
</script>

<div class="border border-primary p-4 bg-base-100 rounded-sm">
    <div class="mb-4">
        <img src="{`${apiUrl}/assets/${selected?.id}`}"
             alt="Preview"
             class="w-full object-contain rounded border"
        />
    </div>

    {#if fileMetadata}
        <div class="text-sm divide-y">
            {#each Object.entries(fileMetadata) as [key, value]}
                <div class="flex justify-between py-2">
                    <span class="font-semibold">{key}</span>
                    <span>{value}</span>
                </div>
            {/each}
        </div>
    {:else}
        <p class="text-gray-500">Lade Metadaten...</p>
    {/if}
</div>