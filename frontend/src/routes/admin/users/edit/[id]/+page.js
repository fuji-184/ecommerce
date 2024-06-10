export async function load({fetch, params}){
  
  try {
    let res = await fetch(`${import.meta.env.VITE_BACKEND_URL}/api/users/${params.id}`, {
        method: "get",
        headers: {
          "content-type": "application/json"
        }
      })
      
      res = await res.json()
      
      return {
        user: res
      }
  } catch (err){
    console.log(err)
  }
  
}