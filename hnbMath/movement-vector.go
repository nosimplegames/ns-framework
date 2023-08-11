package hnbMath

import "math"

type MovementVector struct {
	Speed    float64
	Rotation float64
}

func (vector MovementVector) Calculate() Vector {
	movementVector := Vector{
		X: math.Cos(vector.Rotation) * vector.Speed,
		Y: math.Sin(vector.Rotation) * vector.Speed,
	}

	return movementVector
}
