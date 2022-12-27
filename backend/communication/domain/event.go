package eventingDomain

import "reflect"

type BaseEvent struct {
	Type string
}

func EventType(eventType any) BaseEvent {
	return BaseEvent{Type: reflect.TypeOf(eventType).String()}
}
