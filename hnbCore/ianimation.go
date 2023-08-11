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

	IsAlive() bool
	IsRunning() bool
	HasFinished() bool
	Pause()
	Resume()
	Stop()

	Copy(IAnimationTarget) IAnimation
}
