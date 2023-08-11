package hnbEvents

type EventListener struct {
	EventType string
	Callback  EventCallback
}
