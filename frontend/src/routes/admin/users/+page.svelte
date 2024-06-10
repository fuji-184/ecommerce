<script>
  import { onMount } from "svelte"
  import Tabel from "$lib/components/Tabel.svelte"
  
  let kolom = [
    "nama",
    "username"
  ]
  let data = []
  
  async function getUsers(){
    try {
      let res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/users`, {
        method: "get",
        headers: {
          "content-type": "application/json"
        }
      })
      
      data = await res.json()
    } catch (err){
      console.log("error saat mengambil data: ", err)
    }
  }

  onMount(() => {
    getUsers()
  })
</script>

<div class="mt-4">
  <div class="p-4 bg-purple-100">
    <a class="text-md font-bold p-2 rounded-lg bg-purple-950 text-white" href="/admin/users/add">Add</a>
  </div>
  {#if data}
    <Tabel {kolom} {data} />
  {:else}
  Loading
  {/if}
</div>