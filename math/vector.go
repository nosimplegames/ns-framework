package math

type Vector struct {
	X float64
	Y float64
}

func (vector Vector) By(factor float64) Vector {
	return Vector{
		X: vector.X * factor,
		Y: vector.Y * factor,
	}
}

func (vector Vector) ByVector(another Vector) Vector {
	return Vector{
		X: vector.X * another.X,
		Y: vector.Y * another.Y,
	}
}

func (vector Vector) DivideV(another Vector) Vector {
	return Vector{
		X: vector.X / another.X,
		Y: vector.Y / another.Y,
	}
}

func (vector Vector) Minus(another Vector) Vector {
	return Vector{
		X: vector.X - another.X,
		Y: vector.Y - another.Y,
	}
}

func (vector Vector) Add(another Vector) Vector {
	return Vector{
		X: vector.X + another.X,
		Y: vector.Y + another.Y,
	}
}
