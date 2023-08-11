package hnbCore

import (
	"image/color"

	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type Colorable struct {
	ColorM hnbMath.ColorM
}

func (colorable *Colorable) SetAlpha(alpha float64) {
	colorable.ColorM.SetElement(3, 3, alpha)
}
func (colorable *Colorable) SetRed(alpha float64) {
	colorable.ColorM.SetElement(0, 0, alpha)
}
func (colorable *Colorable) SetGreen(alpha float64) {
	colorable.ColorM.SetElement(1, 1, alpha)
}
func (colorable *Colorable) SetBlue(alpha float64) {
	colorable.ColorM.SetElement(2, 2, alpha)
}

func (colorable *Colorable) SetColor(color color.Color) {
	r, g, b, a := hnbRender.ExtractFloatColorComponents(color)
	colorable.SetRed(r)
	colorable.SetGreen(g)
	colorable.SetBlue(b)
	colorable.SetAlpha(a)
}
