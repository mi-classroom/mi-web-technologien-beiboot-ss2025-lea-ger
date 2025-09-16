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

  let showCreateModal = false;
  let newFolderName = '';
  let creating = false;
  let errorMsg = '';

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

  function openCreateModal() {
    newFolderName = '';
    errorMsg = '';
    showCreateModal = true;
  }

  async function submitCreateFolder() {
    if (!newFolderName.trim()) {
      errorMsg = 'Bitte einen Ordnernamen eingeben.';
      return;
    }
    creating = true;
    errorMsg = '';
    try {
      const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/folders`, {
        method: 'POST',
        headers: {'Content-Type': 'application/json'},
        body: JSON.stringify({name: newFolderName.trim(), parent_id: selectedFolder?.id || null})
      });
      if (response.ok) {
        await loadFolders();
        showCreateModal = false;
        newFolderName = '';
      } else {
        errorMsg = 'Fehler beim Erstellen des Ordners';
      }
    } catch (error) {
      console.error('Fehler beim Erstellen des Ordners', error);
      errorMsg = 'Fehler beim Erstellen des Ordners';
    } finally {
      creating = false;
    }
  }

  $: selectedFolderId = selectedFolder?.id || 0;

  loadFolders();
</script>

<h2 class="text-2xl font-bold mb-6">Ordner</h2>

<button class="btn btn-primary btn-sm mb-6" on:click={openCreateModal}>
    <span class="material-symbols-outlined">create_new_folder</span>
    Ordner erstellen
</button>

{#if showCreateModal}
  <div class="modal modal-open">
    <div class="modal-box">
      <h3 class="font-bold text-lg mb-2">Neuen Ordner erstellen</h3>
      <label class="block mb-2">Ordnername</label>
      <input
        type="text"
        class="input input-bordered w-full"
        placeholder="Name des Ordners"
        bind:value={newFolderName}
        on:keydown={(e) => { if (e.key === 'Enter') submitCreateFolder(); }}
      />
      {#if errorMsg}
        <div class="text-error mt-2">{errorMsg}</div>
      {/if}
      <div class="modal-action">
        <button class="btn" on:click={() => { if (!creating) { showCreateModal = false; errorMsg=''; } }} disabled={creating}>Abbrechen</button>
        <button class="btn btn-primary" on:click={submitCreateFolder} disabled={creating} class:loading={creating}>
          {creating ? 'Erstelle...' : 'Erstellen'}
        </button>
      </div>
    </div>
  </div>
{/if}

{#key rerenderTree}
    <TreeView
            tree={rootFolder}
            selected={selectedFolder}
            onSelect={(folder) => {
                onSelect(folder)
            }}
            folderDeleted={loadFolders}
    />
{/key}