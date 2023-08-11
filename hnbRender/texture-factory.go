package hnbRender

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type TextureFactory struct {
	TargetSize hnbMath.Vector
}

func (factory TextureFactory) Create() Texture {
	texture, _ := ebiten.NewImage(
		int(factory.TargetSize.X),
		int(factory.TargetSize.Y),
		ebiten.FilterDefault,
	)

	return texture
}
