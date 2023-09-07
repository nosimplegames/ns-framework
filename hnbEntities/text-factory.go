package hnbEntities

import (
	"image/color"

	"golang.org/x/image/font"
)

type TextFactory struct {
	Text     string
	FontFace font.Face
	Color    color.Color
}

func (factory TextFactory) Create() *Text {
	text := &Text{}

	text.text = factory.Text
	text.fontFace = factory.FontFace
	text.SetColor(factory.Color)

	return text
}
