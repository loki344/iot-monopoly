import type { Property } from "$lib/model/PropertyDetail";

export async function load({ params, url }) {

    let buyerId = url.searchParams.get('buyerId');

    let property  : Property = {
        Id: params.id ,
        Name: url.searchParams.get('name'),
        FinancialDetails: {
            PropertyPrice: url.searchParams.get('propertyPrice'),
            HousePrice: url.searchParams.get('housePrice'),
            HotelPrice: url.searchParams.get('hotelPrice'),
            Revenue: {
                Normal: url.searchParams.get('revenueNormal'),
                OneHouse: url.searchParams.get('revenueOneHouse'),
                TwoHouses: url.searchParams.get('revenueTwoHouses'),
                ThreeHouses: url.searchParams.get('revenueThreeHouses'),
                FourHouses: url.searchParams.get('revenueFourHouses'),
                Hotel: url.searchParams.get('revenueHotel'),
            }
        }
    }

    return { property: property,
            buyerId: buyerId};
}
