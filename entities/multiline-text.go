package entities

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"golang.org/x/image/font"
)

type MultilineText struct {
	core.Entity
	core.Colorable

	Lines      []string
	LineHeight float64
	FontFace   font.Face
}

func (text MultilineText) Draw(target render.RenderTarget, transform math.Transform) {
	render.MultilineText{
		Lines:      text.Lines,
		Target:     target,
		Position:   text.GetPosition(),
		ColorM:     text.ColorM,
		LineHeight: text.LineHeight,
		FontFace:   text.FontFace,
	}.Render()
}
