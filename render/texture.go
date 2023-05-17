package render

import "github.com/nosimplegames/ns-framework/math"

type Texture = RenderTarget

func GetTextureSize(texture Texture) math.Vector {
	textureBounds := texture.Bounds()

	return math.Vector{
		X: float64(textureBounds.Max.X - textureBounds.Min.X),
		Y: float64(textureBounds.Max.Y - textureBounds.Min.Y),
	}
}
