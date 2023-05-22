package animations

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/utils"
)

type SpriteAnimation struct {
	Animation

	texture           render.Texture
	currentFrameIndex int
	frames            []math.Box
	frameDuration     float64
	currentFrameTime  float64
	target            ISprite
}

func (animation *SpriteAnimation) Update() {
	if !animation.IsRunning() {
		return
	}

	animation.currentFrameTime += utils.FrameTime
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

func (animation SpriteAnimation) GetCurrentRect() *math.Box {
	var rect *math.Box = &math.Box{}

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

func (animation SpriteAnimation) Copy(target utils.Any) core.IAnimation {
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
	Texture       render.Texture
	FrameDuration float64
	FrameSize     math.Vector
	FrameCount    int
	LoopCount     int
}

func (factory SpriteAnimationFactory) Create(target ISprite) core.IAnimation {
	animation := &SpriteAnimation{}

	frames := render.TextureFrameExtractor{
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
