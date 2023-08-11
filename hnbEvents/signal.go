package hnbEvents

import "github.com/nosimplegames/ns-framework/hnbUtils"

type SignalData interface{}
type SignalTCallback func(data SignalData)
type SignalCallback = hnbUtils.Callback

type Signal struct {
	Callbacks []SignalTCallback
}

func (signal *Signal) AddTCallback(callback SignalTCallback) {
	signal.Callbacks = append(signal.Callbacks, callback)
}

func (signal *Signal) AddCallback(callback SignalCallback) {
	signal.Callbacks = append(signal.Callbacks, func(_ SignalData) {
		callback()
	})
}

func (signal *Signal) TFire(value SignalData) {
	for _, callback := range signal.Callbacks {
		callback(value)
	}
}

func (signal *Signal) Fire() {
	signal.TFire(nil)
}
