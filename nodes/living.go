package nodes

import "github.com/nosimplegames/ns-framework/events"

type Living struct {
	IsDead   bool
	IsPaused bool
	OnDie    events.Signal
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
