<script lang="ts">
    import {getAllTags} from "@/utils.js";
    import {onMount} from "svelte";
    import type {IPTCTag} from "@/utils.js";

    export let label: string;
    export let onSelect: (tag: string) => void;
    let iptcTags: IPTCTag[] = [];
    let selectedTag: string | null = null;

    async function loadTags() {
        iptcTags = await getAllTags();
    }

    onMount(loadTags);
</script>

<div class="form-control flex gap-2 items-center mt-4">
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