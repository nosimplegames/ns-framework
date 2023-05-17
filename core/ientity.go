package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/nodes"
)

type IEntity interface {
	nodes.ILiving
	IDrawable
	math.ITransformable
	nodes.INode[IEntity]

	HandleInput()
	Update()

	GetAncestorsTransform() math.Transform
	GetSize() math.Vector
}
