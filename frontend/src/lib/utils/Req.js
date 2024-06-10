export async function req(path, method, mode = null) {
  let headers = {}

  if (mode === "json") {
    headers = { ...headers, "content-type": "application/json" }
  }

  let settings = {
    method: method,
    headers: headers
  };

  try {
    const response = await fetch(`${import.meta.env.VITE_BACKEND_URL}/${path}`, settings)
    const data = await response.json()

    return response
  } catch (error) {
    console.error("Error mengambil data:", error)
  }
}
