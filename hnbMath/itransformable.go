package hnbMath

type ITransformable interface {
	GetTransform() Transform
	GetPosition() Vector
	SetPosition(Vector)
	GetPrevPosition() Vector
	GetOrigin() Vector
	SetOrigin(Vector)
	Move(movementVector Vector)
	MoveX(x float64)
	MoveY(y float64)
	Rotate(tetha float64)
	GetRotation() float64
	SetRotation(rotation float64)
	GetScale() Vector
	SetScale(scale Vector)
}
