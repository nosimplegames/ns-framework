package events

type EventListener struct {
	EventType string
	Callback  EventCallback
}
