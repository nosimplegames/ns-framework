package render

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/nosimplegames/ns-framework/math"
)

type Sprite struct {
	Texture   Texture
	Rect      *math.Box
	Target    RenderTarget
	Transform math.Transform
	ColorM    ebiten.ColorM
}

func (sprite Sprite) Render() {
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
