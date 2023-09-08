package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbEvents"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbNodes"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type DrawPolicy int

const (
	DrawBeforeChildren DrawPolicy = iota
	DrawAfterChildren
)

type IEntity interface {
	IDrawable
	hnbPhysics.ICollisionable
	hnbNodes.INode[IEntity]
	hnbEvents.IEventTarget

	HandleInput()
	UpdateFrame()

	GetAncestorsTransform() hnbMath.Transform
	GetSize() hnbMath.Vector
	SetSize(hnbMath.Vector)

	GetDrawPolicy() DrawPolicy
}
