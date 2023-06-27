package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/nodes"
)

type Entity struct {
	math.Transformable
	nodes.Node[IEntity]
	Drawable

	size       math.Vector
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

	parent, hasParent := entity.GetParent()

	if hasParent {
		entityParent := parent.(IEntity)
		parentTransform := entityParent.GetAncestorsTransform()
		transform.Concat(parentTransform)
		transform.Concat(entityParent.GetTransform())
	}

	return transform
}

func (entity *Entity) SetSize(size math.Vector) {
	entity.size = size
}

func (entity Entity) GetSize() math.Vector {
	return entity.size
}

func (entity *Entity) SetDrawPolicy(drawPolicy DrawPolicy) {
	entity.drawPolicy = drawPolicy
}

func (entity Entity) GetDrawPolicy() DrawPolicy {
	return entity.drawPolicy
}
