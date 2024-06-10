<script>
  let nama, harga, stok
  
  async function simpan(){
    try {
    await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/items`, {
        method: "post",
        headers: {
          "content-type": "application/json"
        },
        body: JSON.stringify({
          Nama: nama,
          Harga: harga,
          Stok: stok,
        })
      })
  } catch (err){
    console.log(err)
  }
  }
  
</script>

<div class="w-[100%] p-4 flex flex-col gap-4">
  <div>
    <label class="text-md font-bold mb-2">nama</label>
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="text" placeholder="nama" bind:value={nama} />
  </div>
  
  <div>
    <label class="text-md font-bold mb-2">harga</label>
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="number" placeholder="harga" bind:value={harga} />
  </div>
  
  <div>
    <label class="text-md font-bold mb-2">stok</label>
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="number" placeholder="stok" bind:value={stok} />
  </div>
  
  <button on:click={simpan} class="p-2 bg-purple-950 text-white rounded-lg mt-4">Simpan</button>
</div>