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
	camera.SetSize(viewport.Size)
	camera.SetOrigin(viewport.Size.By(0.5))
}
