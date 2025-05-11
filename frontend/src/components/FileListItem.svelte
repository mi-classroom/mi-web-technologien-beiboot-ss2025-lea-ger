<script lang="ts">
  export let file;
  export let isSelected = false;
  export let click: () => void;
</script>

<div
        on:click={() => click()}
        class="flex items-center p-2 bg-base-100 dark:bg-dark rounded-sm border cursor-pointer shadow-sm hover:shadow transition-all duration-200"
        class:selected={isSelected}
        class:border-primary={isSelected}
>
    <div class="max-w-12 max-h-16 overflow-hidden">
        <img src={`/${file.thumb}`} alt={file.name}
             class="text-xs  object-cover rounded mr-3"/>
    </div>
    <div class="flex-1">
        <p class="text-sm font-semibold truncate">{file.name}</p>
        <p class="text-xs text-gray-400">{file.date}</p>
    </div>
    <div class="flex flex-col md:flex-row items-center ml-2 text-darkest">
        <button class="btn btn-xs btn-ghost material-symbols-outlined"> edit</button>
        <button class="btn btn-xs btn-ghost material-symbols-outlined">delete</button>
        <button class="btn btn-xs btn-ghost material-symbols-outlined">download</button>
    </div>
</div>
