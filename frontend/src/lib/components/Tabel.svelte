<script>
  import Confirm from "$lib/components/Confirm.svelte"

  export let kolom
  export let data
  
  let no = 1
  let openKonfirmasi = false
  let idForHapus 
  
  function hapus(id){
    openKonfirmasi = true
    idForHapus = id
  }
  
  async function ya(){
    await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/users/${idForHapus}`, {
      method: "delete",
      headers: {
        "content-type": "application/json"
      }
    })
    
    openKonfirmasi = false
  }
  
  function tidak(){
    openKonfirmasi = false 
  }
</script>

<table class="w-[100%] bg-purple-100" border=1>
  <thead>
    <tr>
      <th class="bg-gradient-to-br from-blue-950 to-purple-950 text-white text-md p-4">No</th>
      {#each kolom as k}
        <th class="bg-gradient-to-br from-blue-950 to-purple-950 text-white text-md p-4">{k}</th>
      {/each}
      <th class="bg-gradient-to-br from-blue-950 to-purple-950 text-white text-md p-4">operations</th>
    </tr>
  </thead>
  <tbody>
    {#each data as d, index}
      <tr>
        <td class="text-center border border-purple-900 text-sm p-4">{index + 1}</td>
        {#each kolom as k}
          <td class="text-center border border-purple-900 text-sm p-4">{d[k]}</td>
        {/each}
        
        <td class="text-center border border-purple-900 text-sm p-4">
          <div class="flex gap-4 justify-center items-center">
            <svg viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg" height="1.5em" width="1.5em"><path fill="currentColor" d="M9 39h2.2l22.15-22.15-2.2-2.2L9 36.8Zm30.7-24.3-6.4-6.4 2.1-2.1q.85-.85 2.125-.825 1.275.025 2.125.875L41.8 8.4q.85.85.85 2.1t-.85 2.1ZM7.5 42q-.65 0-1.075-.425Q6 41.15 6 40.5v-4.3q0-.3.1-.55.1-.25.35-.5L31.2 10.4l6.4 6.4-24.75 24.75q-.25.25-.5.35-.25.1-.55.1Zm24.75-26.25-1.1-1.1 2.2 2.2Z"/>
              <a href="/admin/users/edit/{d.nama}" />
            </svg>
          
            <svg on:click={hapus(d.id)} viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg" height="1.5em" width="1.5em"><path fill="currentColor" d="M13.05 42q-1.2 0-2.1-.9-.9-.9-.9-2.1V10.5H9.5q-.65 0-1.075-.425Q8 9.65 8 9q0-.65.425-1.075Q8.85 7.5 9.5 7.5h7.9q0-.65.425-1.075Q18.25 6 18.9 6h10.2q.65 0 1.075.425.425.425.425 1.075h7.9q.65 0 1.075.425Q40 8.35 40 9q0 .65-.425 1.075-.425.425-1.075.425h-.55V39q0 1.2-.9 2.1-.9.9-2.1.9Zm0-31.5V39h21.9V10.5Zm5.3 22.7q0 .65.425 1.075.425.425 1.075.425.65 0 1.075-.425.425-.425.425-1.075V16.25q0-.65-.425-1.075-.425-.425-1.075-.425-.65 0-1.075.425-.425.425-.425 1.075Zm8.3 0q0 .65.425 1.075.425.425 1.075.425.65 0 1.075-.425.425-.425.425-1.075V16.25q0-.65-.425-1.075-.425-.425-1.075-.425-.65 0-1.075.425-.425.425-.425 1.075Zm-13.6-22.7V39 10.5Z"/></svg>
          </div>
        </td>
        
      </tr>
    {/each}
  </tbody>
  <tfoot></tfoot>
</table>

{#if openKonfirmasi}
  <Confirm text="Apakah Anda yakin ingin menghapus data ini?" {ya} {tidak} />
{/if}