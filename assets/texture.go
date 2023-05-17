package assets

import (
	"bytes"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/nosimplegames/ns-framework/render"
)

type TextureBytes []byte

func LoadTexture(textureBytes TextureBytes) render.Texture {
	textureFile, _, err := image.Decode(bytes.NewReader(textureBytes))

	if err != nil {
		panic(err)
	}

	texture, err := ebiten.NewImageFromImage(textureFile, ebiten.FilterDefault)

	if err != nil {
		panic(err)
	}

	return texture
}
