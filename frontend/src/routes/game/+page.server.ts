import { BASE_URL, extractData } from '$lib/http/backendClient';

/** @type {import('./$types').PageServerLoad} */
export async function load({ params }) {
	let gameResponse = await fetch(`${BASE_URL}/games/current`, {
		method: 'GET'
	});

	console.log(gameResponse)
	return {
		game: await extractData(gameResponse),
	};
}

