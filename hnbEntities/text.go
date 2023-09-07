package hnbEntities

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
	"golang.org/x/image/font"
)

type Text struct {
	hnbCore.Entity
	hnbCore.Colorable

	text     string
	fontFace font.Face
}

func (text Text) Draw(target hnbRender.RenderTarget, _ hnbMath.Transform) {
	hnbRender.Text{
		Text:     text.text,
		FontFace: text.fontFace,
		ColorM:   text.ColorM,
		Position: text.GetPosition(),
		Target:   target,
	}.Render()
}
