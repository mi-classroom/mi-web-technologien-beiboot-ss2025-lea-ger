<script lang="ts">
  import type {IPTCTag} from '@/utils.js';
  import {scale} from 'svelte/transition';

  export let keyName: string;
  export let value: string | string[];
  export let originalValue: string | undefined;
  export let tag: IPTCTag | undefined;
  export let onRemove: (key: string) => void | undefined;
  export let onEdit: (key: string, value: string) => void;

  const dateFields = ['DateCreated', 'DateTimeCreated', 'DigitalCreationDate', 'DigitalCreationTime', 'ReleaseDate', 'ExpirationDate'];
  const listFields = ['Keywords', 'SubjectReference']

  let newListItem = '';
  let dateInputValue = '';
  let selectValue = typeof value === 'string' ? value : (Array.isArray(value) ? value[0] : '');

  $: if (tag && dateFields.includes(tag.name)) {
    dateInputValue = parseOutDate(value as string);
  }

  $: if (tag?.values) {
    if (tag.type === 'digits' || tag.type === 'int16u' || tag.type === 'int8u' || tag.type === 'int32u') {
      selectValue = parseInt(value);
    } else {
      selectValue = value
    }
  }

  const parseOutDate = (dateStr: string): string => {
    if (!dateStr) return '';
    const match = dateStr.match(/^(\d{4}):(\d{2}):(\d{2}) (\d{2}):(\d{2}):(\d{2})\+\d{2}:\d{2}$/);
    if (!match) return '';
    return `${match[1]}-${match[2]}-${match[3]}T${match[4]}:${match[5]}:${match[6]}`;
  };

  const parseInDate = (input: string): string => {
    if (!input) return '';
    let [date, time] = input.split('T');
    if (!time) return '';
    if (time.length === 5) time += ':00';
    const [year, month, day] = date.split('-');
    return `${year}:${month}:${day} ${time}+00:00`;
  };

  function parseSelectValue(e: Event) {
    const v = (e.target as HTMLSelectElement).value;
    onEdit(keyName, v);
  }

  function addListItem() {
    if (!newListItem.trim()) return;
    const arr = Array.isArray(value) ? [...value] : [];
    arr.push(newListItem.trim());
    onEdit(keyName, arr);
    newListItem = '';
  }

  function removeListItem(idx: number) {
    if (!Array.isArray(value)) return;
    const arr = value.slice();
    arr.splice(idx, 1);
    onEdit(keyName, arr);
  }

  function handleDateInput(e: Event) {
    const input = (e.target as HTMLInputElement).value;
    dateInputValue = input;
    onEdit(keyName, parseInDate(input));
  }
</script>

<div class="flex items-center justify-between py-2 gap-2">
  <span class="font-semibold flex items-center gap-1">
        <div class="indicator">
            {#if originalValue && value !== originalValue}
            <span transition:scale
                  class="indicator-item status status-warning"/>
            {/if}
            <span class="pr-1">{keyName}</span>
            {#if tag?.relevant && tag?.description}
                <div class="tooltip" data-tip={tag.description}>
                    <span class="material-symbols-outlined text-sm">info</span>
                </div>
            {/if}
        </div>
  </span>
    <span class="flex flex-1 text-right justify-end">
    {#if tag && tag.writable}
        {#if dateFields.includes(tag.name)}
            <input type="datetime-local"
                   class="input input-bordered w-full max-w-xs"
                   bind:value={dateInputValue}
                   on:input={handleDateInput}
            />
        {:else if listFields.includes(tag.name)}
            <div class="flex flex-col gap-2 items-end justify-end w-full max-w-xs">
              <ul class="list-disc pl-4 w-full">
                {#if Array.isArray(value)}
                  {#each value as item, idx}
                    <li class="flex items-center justify-between">
                      <span>{item}</span>
                      <button class="btn btn-xs btn-ghost btn-error ml-2" on:click={() => removeListItem(idx)}>
                        <span class="material-symbols-outlined">close</span>
                      </button>
                    </li>
                  {/each}
                {/if}
              </ul>
              <div class="flex gap-2 w-full">
                <input
                        type="text"
                        class="input input-bordered flex-1"
                        placeholder="Neues Element"
                        bind:value={newListItem}
                        on:keydown={(e) => { if (e.key === 'Enter') { addListItem(); } }}
                />
                <button class="btn btn-primary btn-sm" on:click={addListItem}>Hinzufügen</button>
              </div>
            </div>
        {:else if tag.values}
            <select class="select select-bordered w-full max-w-xs"
                    bind:value={selectValue}
                    on:change={parseSelectValue}
            >
              {#each tag.values as opt}
                <option value={parseInt(opt.id)}>{opt.value}</option>
              {/each}
            </select>
        {:else if tag.type === 'digits' || tag.type === 'int16u' || tag.type === 'int8u' || tag.type === 'int32u'}
          <input type="number" class="input input-bordered w-full max-w-xs" bind:value={value}
                 on:input={e => onEdit(keyName, e.target?.value)}/>
        {:else}
          <input type="text" class="input input-bordered w-full max-w-xs" bind:value={value}
                 on:input={e => onEdit(keyName, e.target?.value)}/>
        {/if}
    {:else}
      {#if listFields.includes(tag?.name) && Array.isArray(value)}
        <ul class="list-disc pl-4">
          {#each value as item}
            <li>{item}</li>
          {/each}
        </ul>
      {:else}
          {value}
      {/if}
    {/if}
  </span>
    {#if onRemove}
        <button class="btn btn-sm btn-error" on:click={() => onRemove(keyName)}>
            <span class="material-symbols-outlined">delete</span>
        </button>
    {/if}
</div>
