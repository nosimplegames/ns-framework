package hnbAnimations

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type SpriteAnimationFactory struct {
	Texture          hnbRender.Texture
	FrameDuration    float64
	FrameSize        hnbMath.Vector
	FrameCount       int
	LoopCount        int
	TimeBetweenLoops float64
}

func (factory SpriteAnimationFactory) Create(target ISprite) hnbCore.IAnimation {
	animation := &SpriteAnimation{}

	frames := hnbRender.TextureFrameExtractor{
		Texture:    factory.Texture,
		FrameSize:  factory.FrameSize,
		FrameCount: factory.FrameCount,
	}.Extract()

	animation.target = target
	animation.texture = factory.Texture
	animation.frameDuration = factory.FrameDuration
	animation.frames = frames
	animation.LoopCount = factory.LoopCount
	animation.timeBetweenLoops = factory.TimeBetweenLoops

	return animation
}
