package hnbCore

import "github.com/nosimplegames/ns-framework/hnbNodes"

type IUpdatable interface {
	hnbNodes.ILiving

	IsRunning() bool
	UpdateFrame()
}
