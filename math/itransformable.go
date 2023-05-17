package math

type ITransformable interface {
	SetPosition(Vector)
	GetTransform() Transform
}
