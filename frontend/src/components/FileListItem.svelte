<script lang="ts">
  import type {FileItem} from "@/utils.js";

  export let file: FileItem;
  export let isSelected = false;
  export let isOpen = false;
  export let onEdit: () => void;
  export let onSelect: () => void;

  export let removed: (fileId: number) => void;

  const apiUrl = import.meta.env.VITE_APP_API_URL;

  const date = new Date(file.upload_date).toLocaleDateString()

  async function deleteFile(): Promise<void> {
    try {
      const response = await fetch(`${apiUrl}/assets/${file.id}`, {
        method: "DELETE",
      });

      if (response.ok) {
        removed(file.id);
      }
    } catch (error) {
      console.error(`Error when deleting file: ${file.filepath}`, error);
    }
  }

  function downloadFile(): void {
    const link = document.createElement("a");
    link.href = `${apiUrl}/assets/${file.id}`;
    link.download = file.filepath;
    link.target = "_blank";
    link.click();
  }
</script>

<div
        on:click={() => onSelect()}
        class={["list-row flex items-center space-x-2 rounded-sm border cursor-pointer shadow-sm hover:shadow transition-all duration-200",
            isOpen ? "bg-primary/20" : "bg-base-100  dark:bg-dark",
            isSelected ? "border-primary selected" : ""
        ].join(" ")}
>
    <div class="list-col-wrap">
        <input
                type="checkbox"
                class="checkbox checkbox-primary"
                checked={isSelected}
                on:click|stopPropagation={() => onSelect()}
        />
    </div>
    <div class="list-col-wrap h-28 w-28 overflow-hidden">
        <img src={`${apiUrl}/assets/${file.id}`} alt={file.filepath}
             class="text-xs object-contain bg-neutral w-full h-full rounded"/>
    </div>
    <div class="flex-1">
        <p class="max-w-40 text-sm font-semibold truncate">{file.filepath}</p>
        <p class="max-w-40 text-xs">{date}</p>
    </div>
    <div class="list-col-wrap flex items-center mx-2">
        <button
                class="btn btn-xs btn-ghost material-symbols-outlined"
                on:click|stopPropagation={onEdit}
        > edit
        </button>
        <button
                class="btn btn-xs btn-ghost material-symbols-outlined"
                on:click|stopPropagation={deleteFile}
        >delete
        </button>
        <button
                on:click|stopPropagation={downloadFile}
                class="btn btn-xs btn-ghost material-symbols-outlined">
            download
        </button>
    </div>
</div>
