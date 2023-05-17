package render

import "image/color"

func IntToColor(intColor int) color.Color {
	color := color.RGBA{
		A: 0xff,
		R: uint8((intColor & 0xFF0000) >> 16),
		G: uint8((intColor & 0x00FF00) >> 8),
		B: uint8((intColor & 0x0000FF)),
	}

	return color
}

func ExtractFloatColorComponents(color color.Color) (float64, float64, float64, float64) {
	rI, gI, bI, aI := color.RGBA()

	r := float64(rI&0xFF) / 255.0
	g := float64(gI&0xFF) / 255.0
	b := float64(bI&0xFF) / 255.0
	a := float64(aI&0xFF) / 255.0

	return r, g, b, a
}
