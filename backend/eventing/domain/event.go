package eventingDomain

import "reflect"

type SensorEvent struct {
	PlayerId   string
	FieldIndex int
}

type BaseEvent struct {
	Type string
}

func EventType(eventType any) *BaseEvent {
	return &BaseEvent{Type: reflect.TypeOf(eventType).String()}
}
