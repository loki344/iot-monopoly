
export async function load({ params, url }) {
    let senderId = url.searchParams.get('senderId');
    let receiverId = url.searchParams.get('receiverId');
    let amount = url.searchParams.get('amount');

    return { senderId, receiverId, amount };
}
