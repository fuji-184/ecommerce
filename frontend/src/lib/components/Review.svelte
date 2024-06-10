<script>
  let reviews = [
    {
      nama: "Fuji",
      pesan: "Baguss barangnya"
    },
    {
      nama: "Fuji",
      pesan: "Baguss barangnya"
    }
  ]
  
  let pesan_review_baru
  
  function kirimReview(){
    reviews = [...reviews, {
      nama: "Fuji",
      pesan: pesan_review_baru
    }]
    
    pesan_review_baru = ""
  }
</script>

<div class="mt-8">
  <div class="flex gap-2">
    <input bind:value={pesan_review_baru} class="p-2 rounded-xl border border-gray-300 w-[100%] focus:outline-none" type="text" placeholder="Review...." />
    <button on:click={kirimReview} class="p-2 bg-purple-950 text-white rounded-xl w-[72px] text-sm font-semibold">Kirim</button>
  </div>
</div>

<div class="max-h-[320px] overflow-y-scroll">
  {#each reviews as r}
    <div class="p-4 rounded-xl border-b border-gray-200">
      <h3 class="text-sm font-semibold">{r.nama}</h3>
      <p class="text-sm whitespace-wrap h-auto">{r.pesan}</p>
    </div>
  {/each}
</div>