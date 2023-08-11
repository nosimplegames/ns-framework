package hnbAnimations

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type SpriteAnimation struct {
	Animation

	texture           hnbRender.Texture
	currentFrameIndex int
	frames            []hnbMath.Box
	frameDuration     float64
	currentFrameTime  float64
	target            ISprite
}

func (animation *SpriteAnimation) UpdateFrame() {
	if !animation.IsRunning() {
		return
	}

	animation.currentFrameTime += hnbUtils.FrameTime
	needToMoveToNextFrame := animation.currentFrameTime >= animation.frameDuration

	if needToMoveToNextFrame {
		animation.currentFrameIndex++
		animation.currentFrameTime = 0
	}
}

func (animation SpriteAnimation) HasFinished() bool {
	hasFinished := animation.currentFrameIndex >= len(animation.frames)

	return hasFinished
}

func (animation *SpriteAnimation) Restart() {
	animation.currentFrameTime = 0
	animation.currentFrameIndex = 0

	animation.Animation.Restart()
}

func (animation SpriteAnimation) Apply() {
	animation.target.SetTexture(animation.texture)
	animation.target.SetRect(animation.GetCurrentRect())
}

func (animation SpriteAnimation) GetCurrentRect() *hnbMath.Box {
	var rect *hnbMath.Box = &hnbMath.Box{}

	if animation.IsRunning() {
		rect = &animation.frames[animation.currentFrameIndex]
	}

	return rect
}

func (animation *SpriteAnimation) Stop() {
	animation.target.SetRect(nil)
	animation.target.SetTexture(nil)
	animation.Animation.Stop()
}

func (animation SpriteAnimation) Copy(target hnbCore.IAnimationTarget) hnbCore.IAnimation {
	copy := &SpriteAnimation{
		Animation: Animation{
			LoopCount: animation.LoopCount,
		},
		texture:       animation.texture,
		frames:        animation.frames,
		frameDuration: animation.frameDuration,
		target:        target.(ISprite),
	}

	return copy
}

type SpriteAnimationFactory struct {
	Texture       hnbRender.Texture
	FrameDuration float64
	FrameSize     hnbMath.Vector
	FrameCount    int
	LoopCount     int
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

	return animation
}
