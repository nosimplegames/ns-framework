package core

import (
	"github.com/nosimplegames/ns-framework/events"
	"github.com/nosimplegames/ns-framework/utils"
)

type IAnimationTarget utils.Any

type IAnimation interface {
	events.IEventTarget

	Update()
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
