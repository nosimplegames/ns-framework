package core

import (
	"github.com/nosimplegames/ns-framework/nodes"
	"github.com/nosimplegames/ns-framework/utils"
)

type Timer struct {
	nodes.Living

	Timeouts []ITimeout
}

func (timer *Timer) Update() {
	for _, timeout := range timer.Timeouts {
		timeout.Update()
	}

	for _, timeout := range timer.Timeouts {
		timeout.Exec()
	}

	utils.RemoveDead(&timer.Timeouts)
}

func (timer *Timer) AddTimeout(timeout ITimeout) {
	timer.Timeouts = append(timer.Timeouts, timeout)
}

func (timer *Timer) RemoveTimeout(timeoutId string) {
	for _, timeout := range timer.Timeouts {
		isTimeoutToRemove := timeout.GetId() == timeoutId

		if isTimeoutToRemove {
			timeout.Die()
			return
		}
	}
}

var timer *Timer = nil

func GetTimer() *Timer {
	needToInitTimer := timer == nil

	if needToInitTimer {
		timer = &Timer{}
		GetUpdatables().AddUpdatable(timer)
	}

	return timer
}
