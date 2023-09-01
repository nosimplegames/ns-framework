package hnbEvents

type IEventTarget interface {
	AddEventListener(EventListener)
	DispatchEvent(Event)
	Fire(string) error
}
