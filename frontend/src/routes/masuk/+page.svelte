<script>
  import Formulir from "$lib/components/Formulir.svelte"
  
  let fields = [
    { name: 'username', label: 'username:', type: 'text', placeholder: 'username' },
    { name: 'password', label: 'password:', type: 'password', placeholder: 'password'},
  ]
  
  let fetchSettings = {
    path: "/masuk",
    method: "post"
  }
  
  console.log(localStorage.getItem("token"))
  
</script>

<Formulir {fields} {fetchSettings} />