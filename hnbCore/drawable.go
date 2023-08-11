package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type Drawable struct {
	IsHidden bool
}

func (drawable Drawable) Draw(target hnbRender.RenderTarget, combinedTransform hnbMath.Transform) {
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
