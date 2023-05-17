package math

type Transformable struct {
	position Vector
	origin   Vector
	rotation float64

	wasScaleSet bool
	scale       Vector

	wasTransformCalculated bool
	transform              Transform
}

func (transformable *Transformable) Move(movementVector Vector) {
	transformable.SetPosition(transformable.position.Add(movementVector))
}

func (transformable *Transformable) Rotate(tetha float64) {
	transformable.SetRotation(transformable.rotation + tetha)
}

func (transformable Transformable) GetRotation() float64 {
	return transformable.rotation
}

func (transformable *Transformable) SetRotation(rotation float64) {
	transformable.rotation = rotation
	transformable.wasTransformCalculated = false
}

func (transformable Transformable) GetPosition() Vector {
	return transformable.position
}

func (transformable *Transformable) SetPosition(position Vector) {
	transformable.position = position
	transformable.wasTransformCalculated = false
}

func (transformable Transformable) GetScale() Vector {
	if !transformable.wasScaleSet {
		return Vector{
			X: 1,
			Y: 1,
		}
	}

	return transformable.scale
}

func (transformable *Transformable) SetScale(scale Vector) {
	transformable.scale = scale
	transformable.wasScaleSet = true
	transformable.wasTransformCalculated = false
}

func (transformable Transformable) GetOrigin() Vector {
	return transformable.origin
}

func (transformable *Transformable) SetOrigin(origin Vector) {
	transformable.origin = origin
	transformable.wasTransformCalculated = false
}

func (transformable *Transformable) GetTransform() Transform {
	if !transformable.wasTransformCalculated {
		transformable.calculateTransform()
	}

	return transformable.transform
}

func (transformable *Transformable) calculateTransform() {
	transform := Transform{}

	if transformable.wasScaleSet {
		transform.Scale(transformable.scale.X, transformable.scale.Y)
	}

	transform.Translate(
		-transformable.origin.X,
		-transformable.origin.Y,
	)
	transform.Rotate(transformable.rotation)
	transform.Translate(
		transformable.position.X,
		transformable.position.Y,
	)

	transformable.transform = transform
	transformable.wasTransformCalculated = true
}
