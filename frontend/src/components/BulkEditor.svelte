<script lang="ts">
  import type {IPTCTag} from "@/utils.js";
  import {getAllTags} from '@/utils.js';
  import {onMount} from 'svelte';
  import {slide} from 'svelte/transition';

  export let onClose: () => void;
  let iptcTags: IPTCTag[] = [];
  let selectedTags: string[] = [];

  async function loadTags() {
    iptcTags = await getAllTags();
  }

  function saveChanges() {
    console.log('Ausgewählte Tags:', selectedTags);
    closeEditor();
  }

  function closeEditor() {
    onClose();
  }

  onMount(loadTags);
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
                        <button class="btn btn-primary" on:click={saveChanges}>Speichern</button>
                        <button class="btn" on:click={closeEditor}>Abbrechen</button>
                    </div>
                </div>
                <div class="form-control mt-4">
                    <label class="label">
                        <span class="label-text">Tags hinzufügen oder entfernen</span>
                    </label>
                    <select multiple bind:value={selectedTags} class="select select-bordered">
                        {#each iptcTags as tag}
                            <option value={tag.name}>{tag.name}</option>
                        {/each}
                    </select>
                </div>

            </div>
        </div>
</div>