<script>
  import Card from "$lib/components/Card.svelte"
  import Hero from "$lib/components/Hero.svelte"
  import { fetcher } from "$lib/utils/fetcher.js"
  import { onMount } from "svelte"
  
  let items = []
  let visibleItems = []
  let startIndex = 0
  let endIndex = 8
  let totalItems = items.length
  
  let hasil = fetcher({
    path:"/api/items",
    token: JSON.parse(localStorage.getItem("token"))
  })
  .then(hasil => {
    console.log(hasil)
    items = hasil
    visibleItems = items.slice(startIndex, endIndex)
  })
  
  function loadMore(){
    startIndex += 8
    endIndex += 8
    visibleItems = [...visibleItems, ...items.slice(startIndex, endIndex)]
  }

let konfirmasi = false
</script>

<Hero />

<div class="mt-14 grid place-items-center gap-8 md:grid-cols-4">
  {#each visibleItems as v}
    <Card nama={v.nama} harga={v.harga} href="/item/{v.id}" />
  {/each}
</div>

{#if endIndex + 1 != totalItems}
  <button on:click={loadMore} class="block m-auto mt-10 p-4 rounded-lg bg-purple-900 text-white text-lg font-bold">More</button>
{/if}