package hnbCore

import (
	"math"

	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbNodes"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type Entity struct {
	hnbPhysics.Collisionable
	hnbNodes.Node[IEntity]
	Drawable

	drawPolicy DrawPolicy

	hasLifeSpan bool
	lifeSpan    float64
}

func (entity *Entity) HandleInput() {
}

func (entity *Entity) UpdateFrame() {
	if entity.hasLifeSpan {
		entity.lifeSpan -= hnbUtils.FrameTime
		shouldDie := entity.lifeSpan <= 0.0

		if shouldDie {
			entity.Die()
		}
	}
}

func (entity Entity) GetPosition() hnbMath.Vector {
	fullTransform := entity.GetAncestorsTransform()
	entityPosition := entity.Transformable.GetPosition()
	transformedX, transformedY := fullTransform.Apply(entityPosition.X, entityPosition.Y)

	return hnbMath.Vector{
		X: transformedX,
		Y: transformedY,
	}
}

func (entity Entity) GetAncestorsTransform() hnbMath.Transform {
	transform := hnbMath.Transform{}

	parent, hasParent := entity.GetParent()

	if hasParent {
		entityParent := parent.(IEntity)
		parentTransform := entityParent.GetAncestorsTransform()
		transform.Concat(parentTransform)
		transform.Concat(entityParent.GetTransform())
	}

	return transform
}

func (entity *Entity) SetDrawPolicy(drawPolicy DrawPolicy) {
	entity.drawPolicy = drawPolicy
}

func (entity Entity) GetDrawPolicy() DrawPolicy {
	return entity.drawPolicy
}

func (entity *Entity) SetLifeSpan(lifeSpan float64) {
	entity.hasLifeSpan = true
	entity.lifeSpan = math.Abs(lifeSpan)
}

func (entity *Entity) Die() {
	entity.Collisionable.Die()
	entity.Node.Die()
}
