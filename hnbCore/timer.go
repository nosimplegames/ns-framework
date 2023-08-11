package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbNodes"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type Timer struct {
	hnbNodes.Living

	Timeouts []ITimeout
}

func (timer *Timer) UpdateFrame() {
	for _, timeout := range timer.Timeouts {
		timeout.UpdateFrame()
	}

	for _, timeout := range timer.Timeouts {
		timeout.Exec()
	}

	hnbUtils.RemoveDead(&timer.Timeouts)
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

func AddTimeout(timeout ITimeout) {
	timer := GetTimer()
	timer.Timeouts = append(timer.Timeouts, timeout)
}
