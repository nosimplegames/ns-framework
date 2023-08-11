package core

import "github.com/nosimplegames/ns-framework/nodes"

type ITimeout interface {
	nodes.ILiving

	UpdateFrame()
	Exec()
	GetId() string
}
