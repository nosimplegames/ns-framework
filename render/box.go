package render

import (
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/nosimplegames/ns-framework/math"
)

type Box struct {
	Position math.Vector
	Size     math.Vector
	Target   RenderTarget
}

func (box Box) Render() {
	color := IntToColor(0x00ff00)

	topLeftCorner, bottomRightCorner := math.Box{
		Position: box.Position,
		Size:     box.Size,
	}.Corners()

	ebitenutil.DrawLine(
		box.Target,
		topLeftCorner.X,
		topLeftCorner.Y,
		bottomRightCorner.X,
		topLeftCorner.Y,
		color,
	)
	ebitenutil.DrawLine(
		box.Target,
		bottomRightCorner.X,
		topLeftCorner.Y,
		bottomRightCorner.X,
		bottomRightCorner.Y,
		color,
	)
	ebitenutil.DrawLine(
		box.Target,
		bottomRightCorner.X,
		bottomRightCorner.Y,
		topLeftCorner.X,
		bottomRightCorner.Y,
		color,
	)
	ebitenutil.DrawLine(
		box.Target,
		topLeftCorner.X,
		bottomRightCorner.Y,
		topLeftCorner.X,
		topLeftCorner.Y,
		color,
	)
}
