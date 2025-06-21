<script lang="ts">
    import UploadBox from '@/components/UploadBox.svelte';
    import FileListItem from '@/components/FileListItem.svelte';
    import MetadataPanel from '@/components/MetadataPanel.svelte';
    import type {FileItem} from "@/utils.js";

    let files: FileItem[] = [];
    let selected: FileItem[] = [];
    let opened: FileItem | null = null;

    /* Dummy files */
    files.push(
        {
            id: 1,
            filepath: 'example1.jpg',
            upload_date: '2023-10-01T12:00:00Z',
            metadata: {}
        },
        {
            id: 2,
            filepath: 'example2.jpg',
            upload_date: '2023-10-02T12:00:00Z',
            metadata: {}
        }
    );

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

    function selectFile(file: FileItem): void {
        if (selected.map(it => it.id).includes(file.id)) {
            selected = selected.filter(it => it.id !== file.id);
        } else {
            selected = [...selected, file];
        }
    }

    loadFiles();
</script>

<div class="min-h-screen p-6 font-sans">
  <h2 class="text-2xl font-bold text-primary mb-6">IPTC Editor</h2>

  <div class="flex flex-col flex-wrap xl:flex-nowrap lg:flex-row gap-6">
    <div class="flex-1">
      <UploadBox onFilesReceived={handleFilesReceived}/>
    </div>

    <div class="flex-1">
      {#if selected.length > 0}
        <div class="lg:w-full xl:flex-1">
          <div class="flex items-center justify-between mb-4">
            <h3 class="text-lg font-semibold">Ausgewählte Dateien ({selected.length})</h3>
            <button
                class="btn btn-primary btn-sm"
                on:click={() => {
                selected = [];
              }}
            >
              Alle abwählen
            </button>
          </div>
        </div>
      {/if}

      {#if files.length > 0}
        <div class="list space-y-3 overflow-y-auto max-h-[700px]">
          {#each files as file (file.id)}
            <FileListItem
                {file}
                isSelected={selected.map(it => it.id).includes(file.id)}
                onSelect={() => selectFile(file)}
                onEdit={() => (opened = file)}
                removed={(fileId) => {
              files = files.filter((f) => f.id !== fileId);
              if (opened?.id === fileId) {
                opened = null;
              }
            }}
            />
          {/each}
        </div>
      {/if}
    </div>

    {#if selected}
      <div class="lg:w-full xl:flex-1">
        <MetadataPanel {selected}/>
      </div>
    {/if}
  </div>
</div>
