package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbNodes"
)

type DrawPolicy int

const (
	DrawBeforeChildren DrawPolicy = iota
	DrawAfterChildren
)

type IEntity interface {
	IDrawable
	hnbMath.ITransformable
	hnbNodes.INode[IEntity]

	HandleInput()
	UpdateFrame()

	GetAncestorsTransform() hnbMath.Transform
	GetSize() hnbMath.Vector
	SetSize(hnbMath.Vector)

	GetDrawPolicy() DrawPolicy
}
