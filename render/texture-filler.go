package render

import "github.com/nosimplegames/ns-framework/math"

type TextureFiller struct {
	Target  RenderTarget
	Texture Texture

	Columns          int
	Rows             int
	StartingPosition math.Vector
}

func (filler TextureFiller) Render() {
	textureSize := GetTextureSize(filler.Texture)

	for row := 0; row < filler.Rows; row++ {
		textureY := filler.StartingPosition.Y + float64(row)*textureSize.Y

		for column := 0; column < filler.Columns; column++ {
			textureX := filler.StartingPosition.X + float64(column)*textureSize.X

			position := math.Vector{
				X: textureX,
				Y: textureY,
			}

			transformable := math.Transformable{}
			transformable.SetPosition(position)
			transform := transformable.GetTransform()

			Sprite{
				Target:    filler.Target,
				Texture:   filler.Texture,
				Transform: transform,
			}.Render()
		}
	}
}
