<script lang="ts">
  import type {Folder} from "@/utils.js";
  import TreeView from "@/components/TreeView.svelte";

  export let onSelect: (folder: Folder) => void;
  export let selectedFolder: Folder | null = null;

  let rerenderTree = 0;

  let rootFolder: Folder = {
    id: 0,
    name: '/root',
  };

  async function loadFolders(): Promise<void> {
    try {
      const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/folders`);
      if (response.ok) {
        const folders = await response.json();
        rootFolder.children = folders.filter((folder: Folder) => !!folder.id)
        rerenderTree += 1;
      } else {
        console.error('Fehler beim Laden der Ordner');
      }
    } catch (error) {
      console.error('Fehler beim Laden der Ordner', error);
    }
  }

  async function createFolder() {
    const folderName = prompt('Bitte geben Sie den Namen des neuen Ordners ein:');
    if (folderName) {
      try {
        const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/folders`, {
          method: 'POST',
          headers: {'Content-Type': 'application/json'},
          body: JSON.stringify({name: folderName, parent_id: selectedFolder?.id || null})
        });
        if (response.ok) {
          await loadFolders();
        } else {
          console.error('Fehler beim Erstellen des Ordners');
        }
      } catch (error) {
        console.error('Fehler beim Erstellen des Ordners', error);
      }
    }
  }

  $: selectedFolderId = selectedFolder?.id || 0;

  loadFolders();
</script>

<h2 class="text-2xl font-bold mb-6">Ordner</h2>

<button class="btn btn-primary btn-sm mb-6" on:click={createFolder}>
    <span class="material-symbols-outlined">create_new_folder</span>
    Ordner erstellen
</button>

{#key rerenderTree}
    <TreeView
            tree={rootFolder}
            selected={selectedFolder}
            onSelect={(folder) => {
                onSelect(folder)
            }}
    />
{/key}