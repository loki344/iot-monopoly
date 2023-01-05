package repository

import "iot-monopoly/game/domain"

func FindPropertyByIndex(index int) *domain.PropertyField {
	return GetCurrentGame().Board().GetPropertyByIndex(index)
}
