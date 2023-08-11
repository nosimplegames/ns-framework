package hnbRender

import "github.com/nosimplegames/ns-framework/hnbMath"

type Texture = RenderTarget

func GetTextureSize(texture Texture) hnbMath.Vector {
	textureBounds := texture.Bounds()

	return hnbMath.Vector{
		X: float64(textureBounds.Max.X - textureBounds.Min.X),
		Y: float64(textureBounds.Max.Y - textureBounds.Min.Y),
	}
}
