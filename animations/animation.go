package animations

import "github.com/nosimplegames/ns-framework/events"

type AnimationState int

const (
	AnimationInfiniteLoop = -1
)

const (
	AnimationRunning AnimationState = iota
	AnimationPaused
	AnimationStopped
)

type Animation struct {
	events.EventTarget

	State       AnimationState
	LoopCount   int
	CurrentLoop int
}

func (animation Animation) IsAlive() bool {
	return animation.State != AnimationStopped
}

func (animation Animation) IsRunning() bool {
	return animation.State == AnimationRunning
}

func (animation *Animation) Stop() {
	animation.State = AnimationStopped
	animation.DispatchEvent(events.Event{
		Type: "stopped",
	})
}

func (animation *Animation) Pause() {
	animation.State = AnimationPaused
}

func (animation *Animation) Resume() {
	animation.State = AnimationRunning
}

func (animation Animation) WillReplay() bool {
	willReplay := animation.LoopCount == AnimationInfiniteLoop || animation.LoopCount-animation.CurrentLoop > 0
	return willReplay
}

func (animation *Animation) Restart() {
	animation.CurrentLoop++
}
