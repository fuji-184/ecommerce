<script>
  let nama, username, password, is_admin
  
  async function simpan(){
    try {
    await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/users`, {
        method: "post",
        headers: {
          "content-type": "application/json"
        },
        body: JSON.stringify({
          Nama: nama,
          Username: username,
          Password: password,
          Is_admin: is_admin
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
    <label class="text-md font-bold mb-2">username</label>
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="text" placeholder="username" bind:value={username} />
  </div>
  
  <div>
    <label class="text-md font-bold mb-2">password</label>
    <input class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" type="password" placeholder="password" bind:value={password} />
  </div>
  
  <div>
    <label class="text-md font-bold mb-2">Apakah Admin?</label>
    <select class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm" bind:value={is_admin}>
      <option></option>
      <option value="true">Ya</option>
      <option value="false">Tidak</option>
    </select>
  </div>
  
  <button on:click={simpan} class="p-2 bg-purple-950 text-white rounded-lg mt-4">Simpan</button>
</div>