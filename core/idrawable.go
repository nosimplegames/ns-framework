package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type IDrawable interface {
	IsVisible() bool
	Draw(target render.RenderTarget, combinedTransform math.Transform)
}
