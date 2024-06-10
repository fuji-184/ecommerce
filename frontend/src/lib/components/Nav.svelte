<script>
  function keluar(){
    if (localStorage.removeItem("token")){
      console.log("berhasil keluar")
    }
    if(!localStorage.getItem("token")){
      console.log("berhasil keluar")
    }
  }
  
  let bukaMenu = false
  function klikMenu(){
    bukaMenu = !bukaMenu
  }
</script>

<div class="flex justify-between p-4 bg-gradient-to-br  from-blue-950 to-purple-950 text-white relative md:items-center">
  <h1 class="h1 text-2xl font-bold">Ecommerce</h1>
    <svg on:click={klikMenu} class="text-2xl font-bold md:hidden" viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg" height="1.5em" width="1.5em"><path fill="currentColor" d="M7.5 36q-.65 0-1.075-.425Q6 35.15 6 34.5q0-.65.425-1.075Q6.85 33 7.5 33h33q.65 0 1.075.425Q42 33.85 42 34.5q0 .65-.425 1.075Q41.15 36 40.5 36Zm0-10.5q-.65 0-1.075-.425Q6 24.65 6 24q0-.65.425-1.075Q6.85 22.5 7.5 22.5h33q.65 0 1.075.425Q42 23.35 42 24q0 .65-.425 1.075-.425.425-1.075.425Zm0-10.5q-.65 0-1.075-.425Q6 14.15 6 13.5q0-.65.425-1.075Q6.85 12 7.5 12h33q.65 0 1.075.425Q42 12.85 42 13.5q0 .65-.425 1.075Q41.15 15 40.5 15Z"/></svg>
  <ul class:flex={bukaMenu} class:hidden={!bukaMenu} class="absolute top-[100%] left-0 right-0 z-10 bg-gradient-to-b from-purple-950 to-blue-950 p-4 flex-col gap-4 md:static md:flex md:flex-row md:bg-transparent">
    
      <li class="hover:bg-purple-100 hover:text-black hover:p-2 hover:rounded-lg">
        <a href="/">Home</a>
      </li>
      
      {#if !localStorage.getItem("token")}
        <li class="hover:bg-purple-100 hover:text-black hover:p-2 hover:rounded-lg">
          <a href="/daftar">Daftar</a>
        </li>
        <li class="hover:bg-purple-100 hover:text-black hover:p-2 hover:rounded-lg">
          <a href="/masuk">Login</a>
        </li>
      {/if}

      {#if localStorage.getItem("token")}
        <li class="hover:bg-purple-100 hover:text-black hover:p-2 hover:rounded-lg">
          <button on:click={keluar}>Keluar</button>
        </li>
      {/if}
  </ul>
</div>