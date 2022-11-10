import { BASE_URL, extractData } from '$lib/http/backendClient';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	let response = await fetch(`${BASE_URL}/players`, {
		method: 'GET'
	});
	return {
		players: await extractData(response)
	};
}
