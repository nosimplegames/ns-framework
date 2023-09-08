package hnbEvents

import "github.com/nosimplegames/ns-framework/hnbUtils"

type IEventTarget interface {
	AddEventListener(EventListener)
	DispatchEvent(Event)
	Fire(string)
	On(string, hnbUtils.Callback)
	OnEvent(string, EventCallback)
}
