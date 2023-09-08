package hnbNodes

import "github.com/nosimplegames/ns-framework/hnbEvents"

type ILiving interface {
	hnbEvents.IEventTarget

	IsAlive() bool
	Die()

	IsRunning() bool
	Pause()
	Resume()
}
