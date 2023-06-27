package events

type IEventTarget interface {
	AddEventListener(EventListener)
	DispatchEvent(Event)
}
