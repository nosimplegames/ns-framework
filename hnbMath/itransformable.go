package hnbMath

type ITransformable interface {
	SetPosition(Vector)
	GetTransform() Transform
	SetOrigin(Vector)
}
