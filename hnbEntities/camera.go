package hnbEntities

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type Camera struct {
	Sprite
	RenderingBox hnbMath.Box
}

func (camera *Camera) SetTexture(texture hnbRender.Texture) {
	camera.Texture = texture
	camera.Rect = &hnbMath.Box{
		Position: camera.RenderingBox.Position,
		Size:     camera.RenderingBox.Size,
	}
}

func (camera *Camera) SetViewport(viewport hnbMath.Box) {
	camera.SetPosition(viewport.Position)
	camera.SetSize(viewport.Size)
	camera.SetOrigin(viewport.Size.By(0.5))
}
