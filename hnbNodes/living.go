package hnbNodes

import "github.com/nosimplegames/ns-framework/hnbEvents"

type Living struct {
	IsDead   bool
	IsPaused bool
	OnDie    hnbEvents.Signal
}

func (living Living) IsAlive() bool {
	return !living.IsDead
}

func (living *Living) Die() {
	if living.IsDead {
		return
	}

	living.IsDead = true
	living.OnDie.Fire()
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
