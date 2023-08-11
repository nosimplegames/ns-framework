package hnbMath

import "image"

type Box struct {
	Size     Vector
	Position Vector
}

func (box Box) IsColliding(another Box) bool {
	topLeftCorner, bottomRightCorner := box.Corners()
	anotherTopLeftCorner, anotherBottomRightCorner := another.Corners()

	return topLeftCorner.X <= anotherBottomRightCorner.X &&
		topLeftCorner.Y <= anotherBottomRightCorner.Y &&
		bottomRightCorner.X >= anotherTopLeftCorner.X &&
		bottomRightCorner.Y >= anotherTopLeftCorner.Y
}

func (box Box) Corners() (Vector, Vector) {
	halfSize := Vector{
		X: box.Size.X * 0.5,
		Y: box.Size.Y * 0.5,
	}

	topLeftCorner := Vector{
		X: box.Position.X - halfSize.X,
		Y: box.Position.Y - halfSize.Y,
	}
	bottomRightCorner := Vector{
		X: box.Position.X + halfSize.X,
		Y: box.Position.Y + halfSize.Y,
	}

	return topLeftCorner, bottomRightCorner
}

func (box Box) IntCorners() (int, int, int, int) {
	topLeftCorner, bottomRightCorner := box.Corners()

	return int(topLeftCorner.X),
		int(topLeftCorner.Y),
		int(bottomRightCorner.X),
		int(bottomRightCorner.Y)
}

func (box Box) ImageRect() image.Rectangle {
	return image.Rect(box.IntCorners())
}

func BoxFromCorners(topX, topY, bottomX, bottomY float64) Box {
	size := Vector{
		X: bottomX - topX,
		Y: bottomY - topY,
	}

	return Box{
		Position: Vector{
			X: topX + size.X*0.5,
			Y: topY + size.Y*0.5,
		},
		Size: size,
	}
}
