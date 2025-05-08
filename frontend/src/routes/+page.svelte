<script>
    import UploadBox from '@/components/UploadBox.svelte';
    import FileListItem from '@/components/FileListItem.svelte';
    import MetadataPanel from '@/components/MetadataPanel.svelte';

    let files = [
        { name: 'AT_KHM_GG6905_FR001_2008-08_Overall-s.jpg', date: '03.05.2025', thumb: 'img1.jpg' },
        { name: 'AT_KHM_GG6739_FR004_2017-01_Overall-s.jpg', date: '01.12.2023', thumb: 'img2.jpg' },
        { name: 'CH_SORW_1925-1b_FR006_2008-11_Overall-s.jpg', date: '03.05.2025', thumb: 'img3.jpg' },
        { name: 'CH_SORW_1925-1a_FR007_1998-11_Overall-s.jpg', date: '01.12.2023', thumb: 'img4.jpg' },
        { name: 'DE_BStGS_1416_FR005_2006_Overall-s.jpg', date: '12.12.2023', thumb: 'img5.jpg' },
        { name: 'DE_smbGG_564A_FR010_2013_Overall-s.jpg', date: '03.05.2025', thumb: 'img6.jpg' }
    ];

    let selected = files[0];
</script>

<div class="min-h-screen p-6 font-sans">
  <h2 class="text-2xl font-bold text-primary mb-6">IPTC Editor</h2>

  <div class="flex gap-6">
    <!-- Upload Box -->
    <div class="w-1/4">
      <UploadBox />
    </div>

    <!-- File List -->
    <div class="w-1/4 space-y-3 overflow-y-auto max-h-[700px]">
      {#each files as file (file.name)}
        <FileListItem
            {file}
            isSelected={file.name === selected.name}
            on:click={() => (selected = file)}
        />
      {/each}
    </div>

    <!-- Metadata Viewer -->
    <div class="w-1/2">
      <MetadataPanel {selected} />
    </div>
  </div>
</div>
