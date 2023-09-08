package hnbNodes

import "github.com/nosimplegames/ns-framework/hnbEvents"

type Living struct {
	hnbEvents.EventTarget

	IsDead   bool
	IsPaused bool
}

func (living Living) IsAlive() bool {
	return !living.IsDead
}

func (living *Living) Die() {
	if living.IsDead {
		return
	}

	living.IsDead = true
	living.Fire("die")
}

func (living Living) IsRunning() bool {
	return !living.IsPaused
}

func (living *Living) Pause() {
	living.IsPaused = true
}

func (living *Living) Resume() {
	living.IsPaused = false
}
