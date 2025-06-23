<script lang="ts">
  import UploadBox from '@/components/UploadBox.svelte';
  import FileListItem from '@/components/FileListItem.svelte';
  import MetadataPanel from '@/components/MetadataPanel.svelte';
  import type {FileItem} from "@/utils.js";

  let files: FileItem[] = [];
  let selected: FileItem[] = [];
  let opened: FileItem | null = null;
  let showUploadModal = false;

  /* Dummy files */
  files.push(
    {
      id: 1,
      filepath: 'example1.jpg',
      upload_date: '2023-10-01T12:00:00Z',
      modification_date: '2023-10-01T12:00:00Z'
    },
    {
      id: 2,
      filepath: 'example2.jpg',
      upload_date: '2023-10-02T12:00:00Z',
      modification_date: '2023-10-02T12:00:00Z'
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
        <!-- UploadBox im Modal -->
        {#if showUploadModal}
            <div class="modal modal-open">
                <div class="modal-box relative">
                    <button
                            class="btn btn-sm btn-error absolute top-2 right-2"
                            on:click={() => (showUploadModal = false)}
                    >
                        Schließen
                    </button>
                    <h2 class="text-lg font-bold mb-4">
                        <span class="material-symbols-outlined">add_photo_alternate</span>
                        Dateien hochladen
                    </h2>
                    <UploadBox onFilesReceived={handleFilesReceived}/>
                </div>
            </div>
        {/if}

        <div class="flex-1">
            <div class="lg:w-full xl:flex-1 mb-4">
                <div class="mb-4">
                    <button class="btn btn-primary btn-sm" on:click={() => showUploadModal = true}>
                        <span class="material-symbols-outlined">add_photo_alternate</span>
                        Dateien hochladen
                    </button>
                </div>

                <div class="flex items-center justify-between">
                    <h3 class="text-lg font-semibold">Ausgewählte Dateien ({selected.length})</h3>
                    {#if selected.length > 0}
                        <button
                                class="btn btn-primary btn-sm"
                                on:click={() => {
                selected = [];
              }}
                        >
                            Alle abwählen
                        </button>

                    {/if}
                </div>
            </div>
            {#if files.length > 0}
                <div class="list space-y-3 overflow-y-auto max-h-[700px]">
                    {#each files as file (file.id)}
                        <FileListItem
                                {file}
                                isSelected={selected.map(it => it.id).includes(file.id)}
                                isOpen={opened?.id === file.id}
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

        {#if opened}
            <div class="lg:w-full xl:flex-1">
                <MetadataPanel selected={opened} />
            </div>
        {/if}
    </div>
</div>
