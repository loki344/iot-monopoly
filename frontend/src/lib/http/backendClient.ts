export const BASE_URL = "http://localhost:3000"

export async function extractData(response: Response) {
    const json = await response.json()
    return JSON.stringify(json)
}