package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type Drawable struct {
	IsHidden bool
}

func (drawable Drawable) Draw(target render.RenderTarget, combinedTransform math.Transform) {
}

func (drawable *Drawable) IsVisible() bool {
	return !drawable.IsHidden
}

func (drawable *Drawable) Hide() {
	drawable.IsHidden = true
}

func (drawable *Drawable) Show() {
	drawable.IsHidden = false
}
