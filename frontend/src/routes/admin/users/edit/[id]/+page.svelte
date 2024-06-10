<script>
  export let data
  
  async function simpan(){
    try {
    await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/users`, {
        method: "put",
        headers: {
          "content-type": "application/json"
        },
        body: JSON.stringify({
          Id: data.user[0].id,
          Nama: data.user[0].nama,
          Username: data.user[0].username
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
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="text" placeholder={data.user ? data.user[0].nama : ""} bind:value={data.user[0].nama} />
  </div>
  <div>
    <label class="text-md font-bold mb-2">username</label>
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="text" placeholder={data.user ? data.user[0].username : ""} bind:value={data.user[0].username} />
  </div>
  
  <button on:click={simpan} class="p-2 bg-purple-950 text-white rounded-lg mt-4">Simpan</button>
</div>