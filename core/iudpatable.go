package core

import "github.com/nosimplegames/ns-framework/nodes"

type IUpdatable interface {
	nodes.ILiving

	IsRunning() bool
	UpdateFrame()
}
