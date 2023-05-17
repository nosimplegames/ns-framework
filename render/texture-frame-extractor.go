package render

import "github.com/nosimplegames/ns-framework/math"

type TextureFrameExtractor struct {
	Texture    Texture
	FrameSize  math.Vector
	FrameCount int
}

func (extractor TextureFrameExtractor) Extract() []math.Box {
	frames := []math.Box{}

	textureWidth, textureHeight := extractor.Texture.Size()
	columns := textureWidth / int(extractor.FrameSize.X)
	rows := textureHeight / int(extractor.FrameSize.Y)
	halfFrameSize := extractor.FrameSize.By(0.5)

	framesToExtract := extractor.GetNumberOfFramesToExtract()

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			mayExtractFrame := framesToExtract == -1 || framesToExtract > 0
			if !mayExtractFrame {
				return frames
			}

			frame := math.Box{
				Size: extractor.FrameSize,
				Position: math.Vector{
					X: halfFrameSize.X + float64(column)*extractor.FrameSize.X,
					Y: halfFrameSize.Y + float64(row)*extractor.FrameSize.Y,
				},
			}
			frames = append(frames, frame)
			framesToExtract--
		}
	}

	return frames
}

func (extractor TextureFrameExtractor) GetNumberOfFramesToExtract() int {
	hasFrameLimit := extractor.FrameCount >= 1

	if !hasFrameLimit {
		return -1
	}

	return extractor.FrameCount
}
