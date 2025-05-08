<script>
    export let file;
    export let isSelected = false;
    export let click;
</script>

<div
    on:click={() => click()}
    class="flex items-center p-2 bg-white rounded border cursor-pointer shadow-sm hover:shadow transition-all duration-200"
    class:selected={isSelected}
    class:border-primary={isSelected}
>
  <img src={`/${file.thumb}`} alt={file.name} class="w-12 h-16 object-cover rounded mr-3" />
  <div class="flex-1">
    <p class="text-sm font-semibold truncate">{file.name}</p>
    <p class="text-xs text-gray-400">{file.date}</p>
  </div>
  <div class="flex items-center space-x-2 ml-2 text-gray-500">
    <button class="btn btn-xs btn-ghost">✏️</button>
    <button class="btn btn-xs btn-ghost">🗑️</button>
    <button class="btn btn-xs btn-ghost">⬇️</button>
  </div>
</div>
