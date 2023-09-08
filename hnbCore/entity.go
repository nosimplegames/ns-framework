package hnbCore

import (
	"math"

	"github.com/nosimplegames/ns-framework/hnbEvents"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbNodes"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type Entity struct {
	hnbPhysics.Collisionable
	hnbNodes.Node[IEntity]
	hnbEvents.EventTarget
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
	return entity.TransformPointByFullTransform(entity.Collisionable.GetPosition())
}

func (entity Entity) GetPrevPosition() hnbMath.Vector {
	return entity.TransformPointByFullTransform(entity.Collisionable.GetPrevPosition())
}

func (entity Entity) TransformPointByFullTransform(point hnbMath.Vector) hnbMath.Vector {
	fullTransform := entity.GetAncestorsTransform()
	transformedX, transformedY := fullTransform.Apply(point.X, point.Y)

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

func (entity Entity) GetAABB() hnbMath.Box {
	return hnbMath.Box{
		Position: entity.GetPosition(),
		Size:     entity.GetSize(),
	}
}

func (entity Entity) GetPrevAABB() hnbMath.Box {
	return hnbMath.Box{
		Position: entity.GetPrevPosition(),
		Size:     entity.GetSize(),
	}
}
