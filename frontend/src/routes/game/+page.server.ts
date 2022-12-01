import { BASE_URL, extractData } from '$lib/http/backendClient';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	let gameDataResponse = await fetch(`${BASE_URL}/players`, {
		method: 'GET'
	});
	let accountResponse = await fetch(`${BASE_URL}/accounts`, {
		method: 'GET'
	});

	return {
		gameData: await extractData(gameDataResponse),
		accounts: await extractData(accountResponse)
	};
}

