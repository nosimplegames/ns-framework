package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/nodes"
)

type DrawPolicy int

const (
	DrawBeforeChildren DrawPolicy = iota
	DrawAfterChildren
)

type IEntity interface {
	IDrawable
	math.ITransformable
	nodes.INode[IEntity]

	HandleInput()
	UpdateFrame()

	GetAncestorsTransform() math.Transform
	GetSize() math.Vector
	SetSize(math.Vector)

	GetDrawPolicy() DrawPolicy
}
