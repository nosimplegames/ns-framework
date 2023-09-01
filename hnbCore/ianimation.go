package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbEvents"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type IAnimationTarget hnbUtils.Any

type IAnimation interface {
	hnbEvents.IEventTarget

	UpdateFrame()
	Apply()

	WillReplay() bool
	Restart()

	WillSleep() bool
	StartSleeping()
	MustWakeUp() bool
	KeepSleeping()

	IsAlive() bool
	IsRunning() bool
	IsSleeping() bool
	HasFinished() bool
	Pause()
	Resume()
	Stop()

	Copy(IAnimationTarget) IAnimation
}
