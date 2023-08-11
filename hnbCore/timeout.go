package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbNodes"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type TimeoutCallback func()

type Timeout struct {
	hnbNodes.Living

	Time        float64
	CurrentTime float64
	Callback    TimeoutCallback
	IsLoop      bool
	Id          string
	ForcedDeath bool
}

func (timeout *Timeout) UpdateFrame() {
	timeout.CurrentTime += hnbUtils.FrameTime
}

func (timeout Timeout) IsAlive() bool {
	return !timeout.ForcedDeath && (timeout.CurrentTime < timeout.Time || timeout.IsLoop)
}

func (timeout *Timeout) Die() {
	timeout.ForcedDeath = true
}

func (timeout *Timeout) Exec() {
	mustExecute := timeout.CurrentTime >= timeout.Time
	if !mustExecute {
		return
	}

	timeout.Callback()

	if timeout.IsLoop {
		timeout.Reset()
	}
}

func (timeout *Timeout) Reset() {
	timeout.CurrentTime = 0
}

func (timeout Timeout) GetId() string {
	return timeout.Id
}
