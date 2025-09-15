<script lang="ts">
  import type {IPTCTag} from "@/utils.js";
  import {getAllTags} from "@/utils.js";
  import {onMount} from "svelte";

  export let label: string;
  export let onSelect: (tag: string) => void;
  export let excludedTags: string[] = [];

  let iptcTags: IPTCTag[] = [];
  let selectedTag: string | null = null;

  async function loadTags() {
    iptcTags = (await getAllTags())
      .filter(tag => !excludedTags.includes(tag.name));
  }

  onMount(loadTags);
</script>

<div class="form-control flex gap-2 justify-end items-center py-4">
    <label class="label">
        <span class="label-text">{label}</span>
    </label>
    <select bind:value={selectedTag} class="select select-bordered">
        {#each iptcTags as tag}
            <option value={tag.name}>{tag.name}</option>
        {/each}
    </select>
    <button class="btn btn-primary" on:click={() => {
        if (selectedTag) {
            onSelect(selectedTag);
            selectedTag = null;
        }
    }}>
        Hinzufügen
    </button>
</div>