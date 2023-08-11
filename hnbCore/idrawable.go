package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type IDrawable interface {
	IsVisible() bool
	Hide()
	Show()
	Draw(target hnbRender.RenderTarget, combinedTransform hnbMath.Transform)
}
