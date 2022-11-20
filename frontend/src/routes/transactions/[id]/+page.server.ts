
export async function load({ params, url }) {
    let senderId = url.searchParams.get('senderId');
    let recipientId = url.searchParams.get('recipientId');
    let amount = url.searchParams.get('amount');

    return { senderId, recipientId, amount };
}
