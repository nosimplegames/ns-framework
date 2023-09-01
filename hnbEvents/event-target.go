package hnbEvents

import (
	"errors"
	"fmt"
)

type EventListeners = map[string][]EventCallback

type EventTarget struct {
	Listeners EventListeners
	creators  []EventCreator
}

func (target *EventTarget) InitListeners() {
	listenersAlreadyInit := target.Listeners != nil

	if !listenersAlreadyInit {
		target.Listeners = EventListeners{}
	}
}

func (target *EventTarget) AddEventListener(listener EventListener) {
	target.InitListeners()

	listenersByType, exists := target.Listeners[listener.EventType]

	if !exists {
		listenersByType = []EventCallback{}
	}

	listenersByType = append(listenersByType, listener.Callback)
	target.Listeners[listener.EventType] = listenersByType
}

func (target *EventTarget) DispatchEvent(event Event) {
	target.InitListeners()

	listenersByType, exists := target.Listeners[event.Type]

	if !exists {
		return
	}

	for _, listener := range listenersByType {
		listener(event)
	}
}

func (target *EventTarget) AddEventCreator(creator EventCreator) {
	target.creators = append(target.creators, creator)
}

func (target EventTarget) Fire(eventType string) error {
	for _, creator := range target.creators {
		canCreatorFire := creator.EventType == eventType

		if canCreatorFire {
			event := creator.CreateEvent()
			target.DispatchEvent(event)
			return nil
		}
	}

	return errors.New(fmt.Sprintf("Event Target: event (%s) cannot be fired from creators", eventType))
}
