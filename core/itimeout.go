package core

import "github.com/nosimplegames/ns-framework/nodes"

type ITimeout interface {
	nodes.ILiving

	Update()
	Exec()
	GetId() string
}
