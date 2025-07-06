<script lang="ts">
  import UploadBox from '@/components/UploadBox.svelte';
  import FileListItem from '@/components/FileListItem.svelte';
  import MetadataPanel from '@/components/MetadataPanel.svelte';
  import type {FileItem, Folder} from "@/utils.js";
  import BulkEditor from "@/components/BulkEditor.svelte";
  import FolderPanel from "@/components/FolderPanel.svelte";

  let files: FileItem[] = [];
  let selected: FileItem[] = [];
  let opened: FileItem | null = null;
  let showUploadModal = false;
  let showEditModal = false;
  let showDeleteConfirmation = false;
  let selectedFolder: {id: number, name: string} | null = null;

  $: isSelected = selected.length > 0;
  $: indeterminate = selected.length > 0 && selected.length < files.length;

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
      formData.append('folder_id', selectedFolder?.id || 0);

      try {
        const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/assets`, {
          method: 'POST',
          body: formData
        });

        if (!response.ok) {
          console.error(`Fehler beim Hochladen der Datei: ${file.name}`);
        }
        showUploadModal = false;
      } catch (error) {
        console.error(`Fehler beim Hochladen der Datei: ${file.name}`, error);
      }
    });

    await Promise.all(uploadPromises);

    await loadFiles();
  }

  async function loadFiles(): Promise<void> {
    const folderId = selectedFolder?.id || 0;
    try {
      const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/metadata?folder_id=${folderId}`);
      if (response.ok) {
        files = (await response.json()) || [];
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

  async function deleteSelected() {
    try {
      for (const item of selected) {
        const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/assets/${item.id}`, {
          method: 'DELETE',
          headers: { 'Content-Type': 'application/json' }
        });

        if (!response.ok) {
          console.error(`Fehler beim Löschen der Datei mit ID: ${item.id}`);
        }
      }

      selected = [];
      loadFiles();
    } catch (error) {
      console.error('Fehler beim Löschen der Dateien', error);
    }
    showDeleteConfirmation = false;
  }

  function handleSelectFolder(folder: Folder): void {
    selectedFolder = folder;
    selected = [];
    opened = null;
    loadFiles();
  }

  loadFiles();
</script>

<div class="min-h-screen p-6 font-sans">
    <h2 class="text-2xl font-bold text-primary mb-6">IPTC Editor</h2>

    <div class="flex flex-col flex-wrap xl:flex-nowrap lg:flex-row gap-6">
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

        <div class="w-1/6 min-w-xs mb-8">
            <FolderPanel
                onSelect={(folder) => {
                    handleSelectFolder(folder);
                }}
                selectedFolder={selectedFolder}
            />
        </div>

        <div class="flex-1">
            <div class="lg:w-full xl:flex-1 mb-4">
                <div class="mb-4">
                    <button class="btn btn-primary btn-sm" on:click={() => showUploadModal = true}>
                        <span class="material-symbols-outlined">add_photo_alternate</span>
                        Dateien hochladen
                    </button>
                </div>

                <div class="flex items-center justify-between px-4">
                    <div class="flex gap-4">
                        <label class="label cursor-pointer">
                            <input
                                    id="select-all"
                                    type="checkbox"
                                    class="checkbox checkbox-primary"
                                    bind:checked={isSelected}
                                    bind:indeterminate={indeterminate}
                                    on:change={() => {
                                        if (selected.length === files.length) {
                                            selected = [];
                                        } else {
                                            selected = [...files];
                                        }
                                    }}
                            />
                        </label>
                        <h3 class="text-lg font-semibold">Ausgewählte Dateien ({selected.length})</h3>
                    </div>

                    <div class="flex gap-4 mt-4">
                        <button class="btn btn-error btn-sm" on:click={showDeleteConfirmation = true}
                                disabled={selected.length === 0}>
                            Löschen
                        </button>
                        <button class="btn btn-primary btn-sm" on:click={() => showEditModal = true}
                                disabled={selected.length === 0}>
                            Bearbeiten
                        </button>
                    </div>
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

        {#if showEditModal}
            <BulkEditor bind:show={showEditModal} onClose={() => showEditModal = false} {selected}/>
        {/if}

        {#if showDeleteConfirmation}
            <div class="modal modal-open">
                <div class="modal-box">
                    <h3 class="font-bold text-lg">Löschen bestätigen</h3>
                    <p>Möchten Sie die ausgewählten Dateien wirklich löschen?</p>
                    <div class="modal-action">
                        <button class="btn btn-error" on:click={deleteSelected}>Ja, löschen</button>
                        <button class="btn" on:click={() => showDeleteConfirmation = false}>Abbrechen</button>
                    </div>
                </div>
            </div>
        {/if}

        {#if opened}
            <div class="lg:w-full xl:flex-1">
                <MetadataPanel selected={opened}/>
            </div>
        {/if}
    </div>
</div>
