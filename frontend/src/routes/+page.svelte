<script lang="ts">
    import UploadBox from '@/components/UploadBox.svelte';
    import FileListItem from '@/components/FileListItem.svelte';
    import MetadataPanel from '@/components/MetadataPanel.svelte';
    import type {FileItem} from "@/utils.js";

    let files: FileItem[] = [];
    let selected: FileItem | null = null;

    async function handleFilesReceived(uploadedFiles: File[]): Promise<void> {
      const uploadPromises = uploadedFiles.map(async (file) => {
        const formData = new FormData();
        formData.append('file', file);

        try {
          const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/assets`, {
            method: 'POST',
            body: formData
          });

          if (!response.ok) {
            console.error(`Fehler beim Hochladen der Datei: ${file.name}`);
          }
        } catch (error) {
          console.error(`Fehler beim Hochladen der Datei: ${file.name}`, error);
        }
      });

      await Promise.all(uploadPromises);

      await loadFiles();
    }

    async function loadFiles(): Promise<void> {
      try {
        const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/metadata`);
        if (response.ok) {
          files = await response.json();
        } else {
          console.error('Fehler beim Laden der Dateien');
        }
      } catch (error) {
        console.error('Fehler beim Laden der Dateien', error);
      }
    }

    loadFiles();
</script>

<div class="min-h-screen p-6 font-sans">
  <h2 class="text-2xl font-bold text-primary mb-6">IPTC Editor</h2>

  <div class="flex flex-col flex-wrap xl:flex-nowrap lg:flex-row gap-6">
    <div class="flex-1">
      <UploadBox onFilesReceived={handleFilesReceived} />
    </div>

    <div class="flex-1 space-y-3 overflow-y-auto max-h-[700px]">
      {#each files as file (file.id)}
        <FileListItem
            {file}
            isSelected={file.id === selected?.id}
            click={() => (selected = file)}
            removed={(fileId) => {
              files = files.filter((f) => f.id !== fileId);
              if (selected?.id === fileId) {
                selected = null;
              }
            }}
        />
      {/each}
    </div>

    {#if selected}
      <div class="lg:w-full xl:flex-2">
        <MetadataPanel {selected} />
      </div>
    {/if}
  </div>
</div>
