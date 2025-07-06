<script lang="ts">
  import {slide} from 'svelte/transition';
  import type {Folder} from "@/utils.js";

  export let tree: Folder;
  export let selected: Folder | null;
  export let onSelect: (item: Folder) => void;
  const {name, children} = tree;
  const _expansionState: { [key: string]: boolean } = {};

  let expanded = _expansionState[name] || false;
  const toggleExpansion = () => {
    expanded = _expansionState[name] = !expanded;
  };
  $: arrowDown = expanded;
</script>

<ul>
    <li class="py-1 px-2 flex flex-wrap items-center gap-1"
        class:bg-primary={selected === tree}
    >
    <span class="material-symbols-outlined">
      {arrowDown ? 'folder_open' : 'folder'}
    </span>
        {#if children}
      <span on:click={toggleExpansion} class="flex items-center">
        <span class="material-symbols-outlined">
          {arrowDown ? 'subdirectory_arrow_right' : 'chevron_forward'}
        </span>
        <span
                class:highlight={selected === tree}
                class="cursor-pointer"
                on:click|stopPropagation={() => onSelect(tree)}>
          {name}
        </span>
      </span>
            {#if expanded}
                {#each children as child (child.name)}
                    <div
                            transition:slide
                            class={'w-full bg-base-300 dark:bg-darker'}
                    >
                        <svelte:self
                                tree={child} {selected} {onSelect}/>
                    </div>
                {/each}
            {/if}
        {:else}
      <span
              class="cursor-pointer"
              on:click|stopPropagation={() => onSelect(tree)}
      >
        {name}
      </span>
        {/if}
    </li>
</ul>