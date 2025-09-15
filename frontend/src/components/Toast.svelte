<script lang="ts">
  export let message: string;
  export let type: 'success' | 'error' | 'info' | 'warning' = 'info';
  export let open: boolean = false;
  export let duration: number = 3000; // ms

  let timer: number | undefined;

  $: if (open && duration > 0) {
    clearTimeout(timer);
    timer = setTimeout(() => {
      open = false;
    }, duration) as unknown as number;
  }

  const bgByType: Record<typeof type, string> = {
    success: 'alert-success',
    error: 'alert-error',
    info: 'alert-info',
    warning: 'alert-warning'
  } as const;
</script>

{#if open}
  <div class="toast toast-top toast-end z-50">
    <div class={"alert shadow-lg " + bgByType[type]}>
      <span>{message}</span>
    </div>
  </div>
{/if}
