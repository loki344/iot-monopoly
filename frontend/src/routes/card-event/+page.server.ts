
export async function load({ params, url }) {
    let text = url.searchParams.get('text');
    let title = url.searchParams.get('title');

    return { text, title };
}
