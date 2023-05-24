package entities

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type Sprite struct {
	core.Entity
	core.Colorable

	Rect    *math.Box
	Texture render.Texture
}

func (sprite *Sprite) SetRect(rect *math.Box) {
	sprite.Rect = rect
}

func (sprite *Sprite) SetTexture(texture render.Texture) {
	sprite.Texture = texture
}

func (sprite *Sprite) SetOriginCenter() {
	size := render.GetTextureSize(sprite.Texture)
	sprite.SetOrigin(size.By(0.5))
}

func (sprite *Sprite) UseTextureSizeAsSize() {
	size := render.GetTextureSize(sprite.Texture)
	sprite.Size = size
}

func (sprite Sprite) Draw(target render.RenderTarget, transform math.Transform) {
	render.Sprite{
		Texture:   sprite.Texture,
		Target:    target,
		Transform: transform,
		Rect:      sprite.Rect,
		ColorM:    sprite.ColorM,
	}.Render()
}

type SpriteFactory struct {
	Texture render.Texture
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

	sprite.Size = render.GetTextureSize(factory.Texture)
	sprite.Texture = factory.Texture
	sprite.SetOrigin(sprite.Size.By(0.5))
}
