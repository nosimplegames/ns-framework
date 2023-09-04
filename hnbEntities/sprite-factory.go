package hnbEntities

import "github.com/nosimplegames/ns-framework/hnbRender"

type SpriteFactory struct {
	Texture hnbRender.Texture
}

func (factory SpriteFactory) Create() *Sprite {
	sprite := &Sprite{}

	factory.Init(sprite)

	return sprite
}

func (factory SpriteFactory) Init(sprite *Sprite) {
	hasTexture := factory.Texture != nil

	if !hasTexture {
		return
	}

	textureSize := hnbRender.GetTextureSize(factory.Texture)
	sprite.SetSize(textureSize)
	sprite.Texture = factory.Texture
	sprite.SetOrigin(textureSize.By(0.5))
}
