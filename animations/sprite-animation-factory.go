package animations

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type SpriteAnimationFactory struct {
	Texture       render.Texture
	FrameDuration float64
	FrameSize     math.Vector
	FrameCount    int
}

func (factory SpriteAnimationFactory) Create() SpriteAnimation {
	frames := render.TextureFrameExtractor{
		Texture:    factory.Texture,
		FrameSize:  factory.FrameSize,
		FrameCount: factory.FrameCount,
	}.Extract()

	animation := SpriteAnimation{
		Animation: Animation{
			State: AnimationRunning,
		},

		Texture:       factory.Texture,
		FrameDuration: factory.FrameDuration,
		Frames:        frames,
	}

	return animation
}
