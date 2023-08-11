package hnbEntities

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
	"golang.org/x/image/font"
)

type MultilineText struct {
	hnbCore.Entity
	hnbCore.Colorable

	Lines      []string
	LineHeight float64
	FontFace   font.Face
}

func (text MultilineText) Draw(target hnbRender.RenderTarget, transform hnbMath.Transform) {
	hnbRender.MultilineText{
		Lines:      text.Lines,
		Target:     target,
		Position:   text.GetPosition(),
		ColorM:     text.ColorM,
		LineHeight: text.LineHeight,
		FontFace:   text.FontFace,
	}.Render()
}
