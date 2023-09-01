package hnbEvents

import "github.com/nosimplegames/ns-framework/hnbUtils"

type EventCreator struct {
	EventType string
	Data      hnbUtils.Any
	GetData   hnbUtils.ReturningCallback
}

func (creator EventCreator) CreateEvent() Event {
	eventData := creator.Data

	hasGetDataCallback := creator.GetData != nil
	if hasGetDataCallback {
		eventData = creator.GetData()
	}

	return Event{
		Type: creator.EventType,
		Data: eventData,
	}
}
