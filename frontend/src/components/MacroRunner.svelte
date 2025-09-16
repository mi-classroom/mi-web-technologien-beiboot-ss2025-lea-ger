<script lang="ts">
  import type { FileItem } from '@/utils.js';
  import { applyMacroToItems, listMacros } from '@/macros.ts';
  import Toast from '@/components/Toast.svelte';

  export let open: boolean = false;
  export let items: FileItem[] = [];
  export let onDone: ((res: { success: number; failed: number }) => void) | undefined;

  let macroOptions: { id: string; name: string }[] = [];
  let selectedMacroIndex = 0;
  let running = false;
  let errorMsg = '';
  let progressDone = 0;
  let progressTotal = 0;

  let toastOpen = false;
  let toastMsg = '';
  let toastType: 'success' | 'error' | 'info' | 'warning' = 'info';

  $: if (open) {
    const list = listMacros();
    macroOptions = list.map(m => ({ id: m.id, name: m.name }));
    selectedMacroIndex = 0;
    errorMsg = '';
    progressDone = 0;
    progressTotal = items.length;
  }

  function cancel() {
    if (running) return;
    open = false;
    errorMsg = '';
  }

  async function confirm() {
    if (items.length === 0) { open = false; return; }
    if (!macroOptions.length) { errorMsg = 'Keine Makros vorhanden.'; return; }
    const macroList = listMacros();
    const idx = selectedMacroIndex;
    if (idx < 0 || idx >= macroList.length) {
      errorMsg = 'Ungültige Auswahl.';
      return;
    }
    running = true;
    try {
      const macro = macroList[idx];
      const res = await applyMacroToItems(items, macro, (done, total) => {
        progressDone = done;
        progressTotal = total;
      });

      if (res.failed === 0) {
        toastType = 'success';
        toastMsg = `Makro ausgeführt. Erfolgreich: ${res.success}`;
      } else if (res.success > 0) {
        toastType = 'warning';
        toastMsg = `Makro teilweise erfolgreich. Erfolgreich: ${res.success}, Fehlgeschlagen: ${res.failed}`;
      } else {
        toastType = 'error';
        toastMsg = 'Makro fehlgeschlagen.';
      }
      toastOpen = true;

      onDone?.(res);
      open = false;
    } catch (e) {
      errorMsg = 'Fehler beim Ausführen des Makros.';
      toastType = 'error';
      toastMsg = 'Fehler beim Ausführen des Makros.';
      toastOpen = true;
    } finally {
      running = false;
    }
  }
</script>

{#if open}
  <div class="modal modal-open">
    <div class="modal-box">
      <h3 class="font-bold text-lg mb-2">Makro ausführen</h3>
      {#if macroOptions.length === 0}
        <p>Keine gespeicherten Makros vorhanden.</p>
      {:else}
        <label class="block mb-2">Wähle ein Makro:</label>
        <select class="select select-bordered w-full" bind:value={selectedMacroIndex}>
          {#each macroOptions as m, i}
            <option value={i}>{m.name}</option>
          {/each}
        </select>
        {#if errorMsg}
          <div class="text-error mt-2">{errorMsg}</div>
        {/if}
        {#if running}
          <div class="mt-4">
            <progress class="progress progress-secondary w-full" value={progressDone} max={progressTotal}></progress>
            <div class="text-xs text-right text-gray-500 mt-1">{progressDone} / {progressTotal} verarbeitet</div>
          </div>
        {/if}
      {/if}
      <div class="modal-action">
        <button class="btn" on:click={cancel} disabled={running}>Abbrechen</button>
        <button class="btn btn-secondary" on:click={confirm} disabled={running || macroOptions.length === 0}>
          {running ? 'Wird ausgeführt...' : 'Makro ausführen'}
        </button>
      </div>
    </div>
  </div>
{/if}

<Toast message={toastMsg} type={toastType} bind:open={toastOpen} duration={3000} />
