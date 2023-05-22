package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type IDrawable interface {
	IsVisible() bool
	Hide()
	Show()
	Draw(target render.RenderTarget, combinedTransform math.Transform)
}
