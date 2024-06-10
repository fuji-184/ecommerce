<!-- src/routes/MyForm.svelte -->
<script>
  import { fetcher } from '$lib/utils/fetcher';
  import { onMount } from 'svelte';

  export let fields = []
  export let fetchSettings = {}

  let formData = {};
  let message = '';
  let imagePreview = null;

  function initializeFormData() {
    fields.forEach(field => {
      formData[field.name] = (field.value) ? field.value : null
    });
  }

  onMount(initializeFormData);

  async function handleSubmit(event) {
    fetchSettings["body"] = formData
    console.log(fetchSettings)
    await fetcher(fetchSettings).then(res => {
      console.log(res)
      if (res.token){
        localStorage.setItem("nama", JSON.stringify(res.nama))
        localStorage.setItem("saldo", JSON.stringify(res.saldo))
        localStorage.setItem("token", JSON.stringify(res.token))
      }
    })
  }

  function handleFileChange(event) {
    const file = event.target.files[0];
    formData[field.name] = file;
    if (file.type.startsWith('image/')) {
      const reader = new FileReader();
      reader.onload = () => {
        imagePreview = reader.result;
      };
      reader.readAsDataURL(file);
    }
  }
</script>

<form on:submit={handleSubmit} class="w-[100%] md:w-[50%] md:m-auto p-4 flex flex-col gap-4">
  {#each fields as field}
  
    <div>
      {#if field.type != 'hidden'}
        <label class="text-md font-bold mb-2">{field.label}</label>
      {/if}

      {#if field.type === 'text'}
        <input type="text" bind:value={formData[field.name]} placeholder={field.placeholder} class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none"/>
      {:else if field.type === 'number'}
        <input type="number" bind:value={formData[field.name]} placeholder={field.placeholder} class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none"/>
      {:else if field.type === 'hidden'}
        <input type="hidden" bind:value={formData[field.name]} />
      {:else if field.type === 'email'}
        <input type="email" bind:value={formData[field.name]} placeholder={field.placeholder} class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" />
      {:else if field.type === 'password'}
        <input type="password" bind:value={formData[field.name]} placeholder={field.placeholder} class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" />
      {:else if field.type === 'file'}
        <div>
          <input type="file" on:change={handleFileChange} class="w-[100%] bg-purple-100 p-2 rounded-lg text-sm focus:outline-none" />
          {#if formData[field.name] && formData[field.name].type.startsWith('image/')}
            <img src={imagePreview} alt="Preview" style="max-width: 200px;" />
          {/if}
        </div>
      {/if}
    </div>
  {/each}

  <button type="submit" class="p-2 bg-purple-950 text-white rounded-lg mt-4">Submit</button>
</form>
