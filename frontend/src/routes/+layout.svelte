<script>
  import "../app.css";
  
  import Nav from "$lib/components/Nav.svelte"
  import Foot from "$lib/components/Foot.svelte"
  
  import eruda from "eruda"
  eruda.init()
</script>

<Nav />
<slot />
<Foot />