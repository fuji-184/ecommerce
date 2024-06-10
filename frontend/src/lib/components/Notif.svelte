<script>
  export let text_notif
  export let ya_notif
</script>

<span class="min-w-[160px] max-w-[40%] absolute left-[50%] top-[50%] translate-x-[-50%] translate-y-[-50%] bg-purple-950 p-4 rounded-lg text-white ">
  <p class="text-md font-semibold mb-4">{text_notif}</p>
  <div class="grid grid-cols-2 gap-2 text-sm font-semibold">
    <button on:click={ya_notif} class="p-1 rounded-lg bg-purple-100 text-purple-950">Ya</button>
  </div>
</span>