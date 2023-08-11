package hnbEvents

type EventListeners = map[string][]EventCallback

type EventTarget struct {
	Listeners EventListeners
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
