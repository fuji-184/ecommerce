<script>
  import Review from "$lib/components/Review.svelte"
  import Confirm from "$lib/components/Confirm.svelte"
  import Notif from "$lib/components/Notif.svelte"
  import { fetcher } from "$lib/utils/fetcher.js"

  export let data
  
  let item
  
  fetcher({ path: `/api/items/${data.id}` }).then(hasil => {
    console.log(data.id)
    item = hasil
    console.log(item)
  })
  
  let saldo = JSON.parse(localStorage.getItem("saldo"))
  
  let konfirmasi = false
  $: text = `Saldo anda Rp${saldo}, klik ya untuk konfirmasi pembelian produk`
  let notif = false
  let text_notif
  
  function ya_notif(){
    notif = false
  }
  
  function beli(){
    konfirmasi = true
  }
  
  function ya(){
    saldo = saldo - item[0].harga
    localStorage.removeItem("saldo")
    localStorage.setItem("saldo", JSON.stringify(saldo))
    text_notif = `Berhasil, saldo anda saat ini Rp${saldo}`
    saldo = JSON.parse(localStorage.getItem("saldo"))
    konfirmasi = false
    notif = true
  }
  
  function tidak(){
    konfirmasi = false
  }
  
</script>

{#if item}
  <div class="p-8 md:grid md:grid-cols-2">
    <div>
      <img class="w-[100%] rounded-xl" src="{import.meta.env.VITE_BACKEND_URL}/files{item[0].gambar}" />
        <h2 class="text-xl font-bold text-purple-950">{item[0].nama}</h2>
        <span class="mt-2 text-xl font-bold text-purple-950">Rp{item[0].harga}</span>
        <button on:click={beli} class="block w-[240px] mt-8 p-4 rounded-xl bg-purple-950 text-white font-bold">Beli Sekarang</button>
    </div>
    <div>
      <Review />
    </div>
  </div>
{/if}

{#if konfirmasi}
  <Confirm {text} {ya} {tidak} />
{/if}

{#if notif}
  <Notif {text_notif} {ya_notif} />
{/if}