package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/nodes"
)

type Entity struct {
	math.Transformable
	nodes.Node[IEntity]
	nodes.Living
	Drawable

	Size       math.Vector
	drawPolicy DrawPolicy
}

func (entity *Entity) HandleInput() {
}

func (entity *Entity) Update() {
}

func (entity Entity) GetPosition() math.Vector {
	fullTransform := entity.GetAncestorsTransform()
	entityPosition := entity.Transformable.GetPosition()
	transformedX, transformedY := fullTransform.Apply(entityPosition.X, entityPosition.Y)

	return math.Vector{
		X: transformedX,
		Y: transformedY,
	}
}

func (entity Entity) GetAncestorsTransform() math.Transform {
	transform := math.Transform{}

	parent := entity.GetParent()
	hasParent := parent != nil

	if hasParent {
		parentTransform := parent.GetAncestorsTransform()
		transform.Concat(parentTransform)
		transform.Concat(parent.GetTransform())
	}

	return transform
}

func (entity Entity) GetSize() math.Vector {
	return entity.Size
}

func (entity *Entity) SetDrawPolicy(drawPolicy DrawPolicy) {
	entity.drawPolicy = drawPolicy
}

func (entity Entity) GetDrawPolicy() DrawPolicy {
	return entity.drawPolicy
}
