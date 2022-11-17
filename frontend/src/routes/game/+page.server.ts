import { BASE_URL, extractData } from '$lib/http/backendClient';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	let playerResponse = await fetch(`${BASE_URL}/players`, {
		method: 'GET'
	});
	let accountResponse = await fetch(`${BASE_URL}/accounts`, {
		method: 'GET'
	});
	return {
		players: await extractData(playerResponse),
		accounts: await extractData(accountResponse)
	};
}

