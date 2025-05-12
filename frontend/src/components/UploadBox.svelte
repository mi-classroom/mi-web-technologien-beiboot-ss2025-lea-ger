<script lang="ts">
  export let onFilesReceived: (files: File[]) => void;

  let files: File[] = [];

  function handleFileChange(event: Event): void {
    const target = event.target as HTMLInputElement;
    if (target?.files) {
      files = Array.from(target.files);
      if (onFilesReceived) {
        onFilesReceived(files);
      }
    }
  }

  function handleDrop(event: DragEvent): void {
    event.preventDefault();
    if (event.dataTransfer?.files) {
      files = Array.from(event.dataTransfer.files);
      if (onFilesReceived) {
        onFilesReceived(files);
      }
    }
  }

  function handleDragOver(event: DragEvent): void {
    event.preventDefault();
  }

  export function receiveFiles(externalFiles: File[]): void {
    files = externalFiles;
    if (onFilesReceived) {
      onFilesReceived(files);
    }
  }
</script>

<div
        class="bg-base-100 rounded-sm border shadow p-6 h-full flex flex-col justify-center items-center text-center"
        on:drop={handleDrop}
        on:dragover={handleDragOver}
>
    <div class="text-3xl mb-4">⬆️</div>
    <p class="text-gray-600 mb-2">Dateien via Drag & Drop hochladen<br/>oder</p>
    <label class="btn bg-[#2fcacb] text-white hover:bg-[#28b8b9]">
        Dateien auswählen
        <input
                type="file"
                multiple
                class="hidden"
                on:change={handleFileChange}
        />
    </label>
</div>