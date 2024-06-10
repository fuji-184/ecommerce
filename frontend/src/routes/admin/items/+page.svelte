<script>
  import { onMount } from "svelte"
  import Tabel from "$lib/components/Tabel.svelte"
  
  import { fetcher } from "$lib/utils/fetcher.js"
  
  let kolom = [
    "nama",
    "harga",
    "stok"
  ]
  let data = []
  
  let hasil = fetcher({
    path:"/api/items",
    token: JSON.parse(localStorage.getItem("token"))
  })
  .then(hasil => {
    console.log(hasil)
    data = hasil
  })
  
</script>

<div class="mt-4">
  <div class="p-4 bg-purple-100">
    <a class="text-md font-bold p-2 rounded-lg bg-purple-950 text-white" href="/admin/items/add">Add</a>
  </div>
  {#if data}
    <Tabel {kolom} {data} />
  {:else}
  Loading
  {/if}
</div>