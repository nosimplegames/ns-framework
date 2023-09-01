package hnbEntities

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type Sprite struct {
	hnbCore.Entity
	hnbCore.Colorable

	Rect    *hnbMath.Box
	Texture hnbRender.Texture
}

func (sprite *Sprite) SetRect(rect *hnbMath.Box) {
	sprite.Rect = rect
}

func (sprite *Sprite) ResetRect() {
	sprite.SetRect(nil)
}

func (sprite *Sprite) SetTexture(texture hnbRender.Texture) {
	sprite.Texture = texture
}

func (sprite *Sprite) SetOriginCenter() {
	size := hnbRender.GetTextureSize(sprite.Texture)
	sprite.SetOrigin(size.By(0.5))
}

func (sprite *Sprite) UseTextureSizeAsSize() {
	size := hnbRender.GetTextureSize(sprite.Texture)
	sprite.SetSize(size)
}

func (sprite Sprite) Draw(target hnbRender.RenderTarget, transform hnbMath.Transform) {
	hnbRender.Sprite{
		Texture:   sprite.Texture,
		Target:    target,
		Transform: transform,
		Rect:      sprite.Rect,
		ColorM:    sprite.ColorM,
	}.Render()
}

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
