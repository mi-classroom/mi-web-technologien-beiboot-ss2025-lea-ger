<script lang="ts">
  import type {IPTCTag} from '@/utils.js';
  import {scale} from 'svelte/transition';

  export let keyName: string;
  export let value: string;
  export let originalValue: string | undefined;
  export let tag: IPTCTag | undefined;
  export let onRemove: (key: string) => void | undefined;
  export let onEdit: (key: string, value: string) => void;
</script>

<div class="flex items-center justify-between py-2 gap-2">
  <span class="font-semibold flex items-center gap-1">
        <div class="indicator">
            {#if originalValue && value !== originalValue}
            <span transition:scale
                   class="indicator-item status status-warning" />
            {/if}
            <span class="pr-1">{keyName}</span>
        </div>
  </span>
    <span class="flex-1 text-right">
    {#if tag && tag.writable}
      {#if tag.values}
        <select class="select select-bordered select-xs w-full max-w-xs" bind:value={value}
                on:change={e => onEdit(keyName, e.target.value)}>
          {#each tag.values as opt}
            <option value={opt.id}>{opt.value}</option>
          {/each}
        </select>
      {:else if tag.type === 'digits' || tag.type === 'int16u' || tag.type === 'int8u' || tag.type === 'int32u'}
          <input type="number" class="input input-bordered w-full max-w-xs" bind:value={value}
                 on:input={e => onEdit(keyName, e.target.value)}/>
      {:else}
          <input type="text" class="input input-bordered w-full max-w-xs" bind:value={value}
                 on:input={e => onEdit(keyName, e.target.value)}/>
      {/if}
    {:else}
      {value}
    {/if}
  </span>
    {#if onRemove}
        <button class="btn btn-sm btn-error" on:click={() => onRemove(keyName)}>
            <span class="material-symbols-outlined">delete</span>
        </button>
    {/if}
</div>

