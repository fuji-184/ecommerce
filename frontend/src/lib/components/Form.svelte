<script>
  let bindVariable = {
    isi: ""
  }
  
  let settings = [
    {
      text: "",
      type: ""
    }
  ]
  
</script>

<form>
  {#each bindVariable as b}
    <input type={input.type} bind:value={b} placeholder={input.text} />
  {/each}
</form>