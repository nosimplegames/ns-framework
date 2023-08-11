package hnbCore

import "github.com/nosimplegames/ns-framework/hnbNodes"

type ITimeout interface {
	hnbNodes.ILiving

	UpdateFrame()
	Exec()
	GetId() string
}
