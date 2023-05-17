package render

import "golang.org/x/image/font"

func GetFontFaceHeight(fontFace font.Face) float64 {
	return float64(fontFace.Metrics().Height.Round())
}
