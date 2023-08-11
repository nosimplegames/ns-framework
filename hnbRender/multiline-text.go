package hnbRender

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"golang.org/x/image/font"
)

type MultilineText struct {
	Lines         []string
	FontFace      font.Face
	LineHeight    float64
	Target        RenderTarget
	Position      hnbMath.Vector
	TextAlignment TextAlignment
	ColorM        hnbMath.ColorM
}

func (text MultilineText) Render() {
	linesRenderingPositions := text.GetLinesRenderingPositions()

	for lineIndex, line := range text.Lines {
		position := linesRenderingPositions[lineIndex]

		Text{
			Text:     line,
			Target:   text.Target,
			FontFace: text.FontFace,
			Position: position,
			ColorM:   text.ColorM,
		}.Render()
	}
}

func (text MultilineText) GetLinesRenderingPositions() []hnbMath.Vector {
	lineFontHeight := GetFontFaceHeight(text.FontFace)
	lineHeight := lineFontHeight * text.LineHeight
	renderingPosition := text.GetRenderingPosition(lineHeight)
	positions := []hnbMath.Vector{}
	lineTopOffset := (lineHeight - lineFontHeight) / 2.0

	for lineIndex := range text.Lines {
		position := hnbMath.Vector{
			X: renderingPosition.X,
			Y: renderingPosition.Y + lineTopOffset + lineHeight*float64(lineIndex),
		}
		positions = append(positions, position)
	}

	return positions
}

func (text MultilineText) GetRenderingPosition(lineHeight float64) hnbMath.Vector {
	return hnbMath.Vector{
		X: text.Position.X,
		Y: text.Position.Y - lineHeight,
	}
}
