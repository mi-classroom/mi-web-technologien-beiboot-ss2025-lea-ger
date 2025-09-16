<script lang="ts">
  import {hasMacros} from "@/macros.ts";
  import {type FileItem, loadFilesInFolder} from "@/utils.js";
  import MacroRunner from "@/components/MacroRunner.svelte";

  export let folderName: string;
  export let folderId: string | number;
  export let onDelete: null | undefined | ((id: string | number) => void);

  $: canDelete = !!onDelete;

  let showDeleteConfirm = false;
  let deleting = false;
  let errorMsg = '';
  let showMacroModal = false;
  let macroFiles: FileItem[] = [];

  async function deleteFolder(id: string | number) {
    deleting = true;
    errorMsg = '';
    try {
      const res = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/folders/${id}`, {method: 'DELETE'});
      if (!res.ok) throw new Error(await res.text());
      if (canDelete) onDelete(id);
    } catch (e) {
      if (e instanceof Error) {
        errorMsg = e.message;
      } else {
        errorMsg = 'Fehler beim Löschen';
      }
    } finally {
      deleting = false;
      showDeleteConfirm = false;
    }
  }

  async function applyMacro() {
    macroFiles = await loadFilesInFolder(folderId);
    console.log(macroFiles);
    showMacroModal = true;
  }

  function handleMacroDone(event: CustomEvent<{ success: number; failed: number }>) {
    macroFiles = [];
    showMacroModal = false;
  }
</script>


<div class="dropdown dropdown-bottom">
    <button aria-label="Aktionen"
            role="button"
            class="btn btn-ghost material-symbols-outlined cursor-pointer ml-1 bg-transparent border-0 p-0 focus:outline-none"
            type="button"
    >
        more_vert
    </button>
    <ul
            class="dropdown-content menu p-2 shadow bg-base-200 rounded-box w-52"
            id={"popover-" + folderId}
            popover
    >
        {#if hasMacros()}
            <li>
                <button class="hover:bg-secondary hover:text-secondary-content flex items-center gap-1"
                        on:click={applyMacro}
                >
                    <span class="material-symbols-outlined">play_arrow</span>
                    Makro ausführen
                </button>
            </li>
        {/if}
        {#if canDelete}
            <li>
                <button class="hover:bg-error hover:text-error-content flex items-center gap-1"
                        on:click={() => { showDeleteConfirm = true; }}>
                    <span class="material-symbols-outlined">delete</span>
                    Ordner löschen
                </button>
            </li>
        {/if}
    </ul>
</div>
<MacroRunner
        bind:open={showMacroModal}
        items={macroFiles}
        onDone={handleMacroDone}
/>

{#if showDeleteConfirm}
    <div class="fixed inset-0 bg-black bg-opacity-40 flex items-center justify-center z-50">
        <div class="bg-base-100 p-6 rounded shadow-lg min-w-[300px]">
            <div class="mb-4">Diesen Ordner wirklich löschen?<br><b>{folderName}</b></div>
            {#if errorMsg}
                <div class="text-error mb-2">{errorMsg}</div>
            {/if}
            <div class="flex gap-2 justify-end">
                <button class="btn btn-ghost" on:click={() => showDeleteConfirm = false} disabled={deleting}>Abbrechen
                </button>
                <button class="btn btn-error" on:click={() => deleteFolder(folderId)} disabled={deleting}>
                    {deleting ? 'Lösche...' : 'Löschen'}
                </button>
            </div>
        </div>
    </div>
{/if}
