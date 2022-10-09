package movement

import (
	"fmt"
	"github.com/vmware/transport-go/bus"
	"github.com/vmware/transport-go/model"
	"iot-monopoly/eventing"
	"iot-monopoly/game_master"
)

func Init() {

	tr := bus.GetBus()
	channel := "external"

	sensorEventHandler, err := tr.ListenRequestStream(channel)
	if err != nil {
		fmt.Println(err)
	}
	sensorEventHandler.Handle(
		func(msg *model.Message) {
			sensorEvent := msg.Payload.(eventing.SensorEvent)
			movePlayer(sensorEvent.PlayerId, sensorEvent.FieldIndex)
		},
		func(err error) {
			fmt.Println(err)
		})
}

func movePlayer(playerId int, fieldIndex int) {

	fmt.Printf("Move player %d to fieldIndex %d\n", playerId, fieldIndex)
	player := game_master.GetPlayer(playerId)
	if (player.Position >= len(game_master.Fields)-6 && player.Position < len(game_master.Fields)) && (fieldIndex >= 0 && fieldIndex <= 5) {
		fmt.Println("Player completed a lap, creating lap finished")
		ts := bus.GetBus()
		handler, err := ts.RequestOnce("lapFinished", LapFinishedEvent{playerId})
		if err != nil {
			//TODO something went wrong
			return
		}
		handler.Fire()
	}

	player.Position = fieldIndex
}

type LapFinishedEvent struct {
	PlayerId int
}
