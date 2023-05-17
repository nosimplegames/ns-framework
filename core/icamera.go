package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type ICamera interface {
	math.ITransformable
	IDrawable

	SetTexture(render.Texture)
}
