async function makeRequest(method: "get" | "post", url: string, body?: any) {
    const res = await fetch(`http://localhost:8081/${url}`, {
        method,
        body: JSON.stringify(body)
    })
    const result = await res.json();
    return result;
}