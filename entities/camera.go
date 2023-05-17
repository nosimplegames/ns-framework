package entities

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type Camera struct {
	Sprite
	RenderingBox math.Box
}

func (camera *Camera) SetTexture(texture render.Texture) {
	camera.Texture = texture
	camera.Rect = &math.Box{
		Position: camera.RenderingBox.Position,
		Size:     camera.RenderingBox.Size,
	}
}

func (camera *Camera) SetViewport(viewport math.Box) {
	camera.SetPosition(viewport.Position)
	camera.Size = viewport.Size
	camera.SetOrigin(camera.Size.By(0.5))
}
