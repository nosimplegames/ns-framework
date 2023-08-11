package hnbRender

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type Sprite struct {
	Texture   Texture
	Rect      *hnbMath.Box
	Target    RenderTarget
	Transform hnbMath.Transform
	ColorM    ebiten.ColorM
}

func (sprite Sprite) Render() {
	canRender := sprite.Texture != nil

	if !canRender {
		return
	}

	drawImageOptions := &ebiten.DrawImageOptions{
		GeoM:   sprite.Transform,
		ColorM: sprite.ColorM,
	}

	sprite.Target.DrawImage(
		sprite.GetRenderingTexture(),
		drawImageOptions,
	)
}

func (sprite Sprite) GetRenderingTexture() Texture {
	mustRenderWholeTexture := sprite.Rect == nil

	if mustRenderWholeTexture {
		return sprite.Texture
	}

	rect := image.Rect(sprite.Rect.IntCorners())
	return sprite.Texture.SubImage(rect).(Texture)
}
