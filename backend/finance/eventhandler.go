package finance

import (
	"context"
	"fmt"
	"github.com/mustafaturan/bus/v3"
	"iot-monopoly/board"
	boardDomain "iot-monopoly/board/domain"
	"iot-monopoly/eventing"
)

func StartEventListeners() {

	startLapFinishedEventHandler()
}

func startLapFinishedEventHandler() {

	eventing.RegisterEventHandler(bus.Handler{
		Handle: func(ctx context.Context, e bus.Event) {
			lapFinishedEvent := e.Data.(boardDomain.LapFinishedEvent)
			fmt.Println("Add money to balance due to lap finished")
			board.GetPlayer(lapFinishedEvent.PlayerId).Balance += 100
		},
		Matcher: string(eventing.LAP_FINISHED),
	})
}
