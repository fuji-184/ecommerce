<script>
  import Formulir from "$lib/components/Formulir.svelte"
  
  export let data
  console.log(data)
  
  let fields = [
    { name: 'nama', label: 'Nama:', type: 'text', placeholder: 'nama', value: data.data[0].nama },
    { name: 'harga', label: 'harga:', type: 'number', placeholder: 'harga', value: data.data[0].harga },
    { name: 'stok', label: 'stok:', type: 'number', placeholder: 'stok', value: data.data[0].stok },
    { name: 'id', type: 'hidden', value: data.data[0].id }
  ]
  
  let fetchSettings = {
    path: "/api/items",
    method: "put"
  }
</script>

<Formulir {fields} {fetchSettings} />