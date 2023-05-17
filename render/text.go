package render

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	ebitenText "github.com/hajimehoshi/ebiten/text"
	"github.com/nosimplegames/ns-framework/math"
	"golang.org/x/image/font"
)

type Text struct {
	Text     string
	FontFace font.Face
	Target   *ebiten.Image
	Position math.Vector
	ColorM   math.ColorM
}

func (text Text) Render() {
	x, y := text.GetRenderingPosition()
	color := text.ColorM.Apply(color.White)

	ebitenText.Draw(
		text.Target,
		text.Text,
		text.FontFace,
		x,
		y,
		color,
	)
}

func (text Text) GetRenderingPosition() (int, int) {
	renderingSize := ebitenText.BoundString(text.FontFace, text.Text)

	drawingPosition := text.Position
	drawingPosition.X -= float64(renderingSize.Max.X) * 0.5
	drawingPosition.Y += float64(-renderingSize.Min.Y)

	return int(drawingPosition.X), int(drawingPosition.Y)
}
