package render

import (
	"math"

	emath "github.com/nosimplegames/ns-framework/math"
)

type TextureRepeater struct {
	Target  RenderTarget
	Texture Texture
	Area    emath.Box
}

func (repeater TextureRepeater) Repeat() {
	columns, rows := repeater.GetColumnsRows()
	topLeftCorner, _ := repeater.Area.Corners()

	for column := 0; column < columns; column++ {
		for row := 0; row < rows; row++ {
			transform := repeater.GetTransform(column, row, topLeftCorner)

			Sprite{
				Target:    repeater.Target,
				Texture:   repeater.Texture,
				Transform: transform,
			}.Render()
		}
	}
}

func (repeater TextureRepeater) GetColumnsRows() (int, int) {
	textureSize := GetTextureSize(repeater.Texture)
	columns := int(math.Ceil(repeater.Area.Size.X / textureSize.X))
	rows := int(math.Ceil(repeater.Area.Size.Y / textureSize.Y))

	return columns, rows
}

func (repeater TextureRepeater) GetTransform(column, row int, startingPosition emath.Vector) emath.Transform {
	transformable := emath.Transformable{}
	textureSize := GetTextureSize(repeater.Texture)

	transformable.SetPosition(startingPosition.Add(
		textureSize.ByVector(
			emath.Vector{
				X: float64(column),
				Y: float64(row),
			},
		),
	))

	return transformable.GetTransform()
}
