package oldanimations

import "github.com/nosimplegames/ns-framework/events"

type AnimationState int

const (
	AnimationRunning AnimationState = iota
	AnimationPaused
	AnimationStopped
)

type Animation struct {
	State  AnimationState
	OnStop events.Signal
}

func (animation Animation) IsAlive() bool {
	return animation.State != AnimationStopped
}

func (animation Animation) IsRunning() bool {
	return animation.State == AnimationRunning
}

func (animation *Animation) Stop() {
	animation.State = AnimationStopped
	animation.OnStop.Fire()
}

func (animation *Animation) Pause() {
	animation.State = AnimationPaused
}

func (animation *Animation) Resume() {
	animation.State = AnimationRunning
}

func (animation *Animation) Die() {
	animation.Stop()
}
