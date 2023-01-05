package domain

import (
	"iot-monopoly/eventing"
	"iot-monopoly/game/domain/events"
)

var defaultFinancialDetails = FinancialDetails{100, 100, 100, Revenue{100, 200, 300, 400, 500, 800}}

//var defaultBasicFields = []domain.BasicField{
//	{domain.BaseFieldInformation{index: "1", name: "Start"}},
//	{domain.BaseFieldInformation{index: "5", name: "Gefaengnis"}},
//	{domain.BaseFieldInformation{index: "9", name: "Frei parken"}},
//}

var defaultProperties = []PropertyField{
	*NewPropertyField("Property purple 1", 2, defaultFinancialDetails),
	*NewPropertyField("Property purple 2", 3, defaultFinancialDetails),
	*NewPropertyField("Property orange 1", 7, defaultFinancialDetails),
	*NewPropertyField("Property orange 2", 8, defaultFinancialDetails),
	*NewPropertyField("Property green 1", 10, defaultFinancialDetails),
	*NewPropertyField("Property green 2", 12, defaultFinancialDetails),
	*NewPropertyField("Property blue 1", 14, defaultFinancialDetails),
	*NewPropertyField("Property blue 2", 16, defaultFinancialDetails),
}
var defaultEventFields = []EventField{
	*NewEventField("Ereignisfeld 1", 4, DRAW_CARD),
	*NewEventField("Ereignisfeld 2", 6, DRAW_CARD),
	*NewEventField("Einkommenssteuer", 11, PAY_TAX),
	*NewEventField("Gehe ins Gefaengnis", 13, GOTO_PRISON),
	*NewEventField("Ereignisfeld 4", 15, DRAW_CARD),
}

var defaultCardStack = []Card{
	*NewCard("You inherited", "You're mentioned in the testament of your aunt. You receive 100 $.", func(player *Player) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_PAYOUT_ACCEPTED, events.NewGameEventWithPayoutAcceptedEvent(player.Account().Id(), 100))
	}),
	*NewCard("Tax bill", "You received a bill for the federal taxes of 200 $", func(player *Player) {
		eventing.FireEvent(eventing.GAME_EVENT_WITH_FEE_ACCEPTED, events.NewGameEventWithFeeAcceptedEvent("Bank", player.Account().Id(), 200))
	}),
}
