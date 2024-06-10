// src/lib/fetcher.js
export async function fetcher({ path = "/", method = "GET", mode = "json", token = "", body = null} = {}) {
    try {
        const url = import.meta.env.VITE_BACKEND_URL + path;

        const headers = {};

        if (token) {
            headers['Authorization'] = `Bearer ${token}`;
        }

        if (mode === "json") {
            headers['Content-Type'] = 'application/json';
        }

        const requestOptions = {
            method,
            headers
        };
        
        console.log(requestOptions)

        if (body !== null) {
            if (mode === "json") {
                requestOptions.body = JSON.stringify(body);
            } else if (mode === "image") {
                requestOptions.body = body;
            }
        }

        const response = await fetch(url, requestOptions);

        if (!response.ok) {
            throw new Error(`HTTP error! status: ${response.status}`);
        }

        if (mode === "json") {
            return await response.json();
        } else if (mode === "image") {
            const blob = await response.blob();
            return URL.createObjectURL(blob);
        } else {
            throw new Error(`Unsupported mode: ${mode}`);
        }
    } catch (error) {
        console.error('Fetch Error:', error);
        throw error;
    }
}
