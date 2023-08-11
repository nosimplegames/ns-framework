package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type ICamera interface {
	hnbMath.ITransformable
	IDrawable

	SetTexture(hnbRender.Texture)
}
