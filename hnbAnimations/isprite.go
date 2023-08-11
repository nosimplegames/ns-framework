package hnbAnimations

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type ISprite interface {
	SetRect(*hnbMath.Box)
	SetTexture(hnbRender.Texture)
	UseTextureSizeAsSize()
}
