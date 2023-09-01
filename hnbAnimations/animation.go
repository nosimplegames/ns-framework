package hnbAnimations

import (
	"github.com/nosimplegames/ns-framework/hnbEvents"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type AnimationState int

const (
	AnimationInfiniteLoop = -1
)

const (
	AnimationRunning AnimationState = iota
	AnimationPaused
	AnimationStopped
	AnimationSleeping
)

type Animation struct {
	hnbEvents.EventTarget

	State               AnimationState
	LoopCount           int
	CurrentLoop         int
	timeBetweenLoops    float64
	currentSleepingTime float64
}

func (animation Animation) IsAlive() bool {
	return animation.State != AnimationStopped
}

func (animation Animation) IsRunning() bool {
	return animation.State == AnimationRunning
}

func (animation Animation) IsSleeping() bool {
	return animation.State == AnimationSleeping
}

func (animation *Animation) Stop() {
	animation.State = AnimationStopped
	animation.DispatchEvent(hnbEvents.Event{
		Type: "stopped",
	})
}

func (animation *Animation) Pause() {
	animation.State = AnimationPaused
}

func (animation *Animation) Resume() {
	animation.State = AnimationRunning
}

func (animation *Animation) StartSleeping() {
	animation.State = AnimationSleeping
	animation.currentSleepingTime = 0

	animation.DispatchEvent(hnbEvents.Event{
		Type: "start-sleeping",
	})
}

func (animation *Animation) KeepSleeping() {
	animation.currentSleepingTime += hnbUtils.FrameTime
}

func (animation Animation) MustWakeUp() bool {
	mustWakeUp := animation.currentSleepingTime >= animation.timeBetweenLoops

	return mustWakeUp
}

func (animation Animation) WillReplay() bool {
	willReplay := animation.LoopCount == AnimationInfiniteLoop || animation.LoopCount-animation.CurrentLoop > 0
	return willReplay
}

func (animation Animation) WillSleep() bool {
	willSleep := animation.timeBetweenLoops > 0.0
	return willSleep
}

func (animation *Animation) Restart() {
	animation.CurrentLoop++
	animation.Resume()
}

func (animation Animation) Copy() Animation {
	return Animation{
		LoopCount:        animation.LoopCount,
		timeBetweenLoops: animation.timeBetweenLoops,
	}
}
