<script lang="ts">
  import type {IPTCTag} from "@/utils.js";
  import {slide} from 'svelte/transition';
  import TagSelector from "@/components/TagSelector.svelte";
  import MetadataField from "@/components/MetadataField.svelte";

  export let onClose: () => void;
  let tags: {[IPTCTag]: string} = {};

  function saveChanges() {
    console.log('Ausgewählte Tags:', selectedTags);
    closeEditor();
  }

  function addTag(tag: string) {
    if (!tags[tag]) {
        tags[tag] = '';
    } else {
        console.warn(`Tag "${tag}" ist bereits ausgewählt.`);
    }
  }

  function closeEditor() {
    onClose();
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
                        <button class="btn btn-primary" on:click={saveChanges}>Speichern</button>
                        <button class="btn" on:click={closeEditor}>Abbrechen</button>
                    </div>
                </div>
                <div class="space-y-2 my-4">
                    {#each Object.entries(tags) as [key, value]}
                        <MetadataField
                                keyName={key}
                                value={value}
                                originalValue={value}
                                tag={tags[key]}
                                onEdit={(k, v) => tags[k] = v}
                        />
                    {/each}
                </div>

                <TagSelector
                        onSelect={addTag}
                />
            </div>
        </div>
</div>