package oldanimations

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/utils"
)

type SpriteAnimation struct {
	Animation

	Texture           render.Texture
	CurrentFrameIndex int
	Frames            []math.Box
	FrameDuration     float64
	CurrentFrameTime  float64
}

func (animation *SpriteAnimation) Update() {
	if !animation.IsRunning() {
		return
	}

	animation.CurrentFrameTime += utils.FrameTime
	needToMoveToNextFrame := animation.CurrentFrameTime >= animation.FrameDuration

	if needToMoveToNextFrame {
		animation.CurrentFrameIndex++
		animation.CurrentFrameTime = 0

		needToStop := animation.CurrentFrameIndex >= len(animation.Frames)

		if needToStop {
			animation.Stop()
		}
	}
}

func (animation SpriteAnimation) GetCurrentRect() *math.Box {
	var rect *math.Box = &math.Box{}

	if animation.IsRunning() {
		rect = &animation.Frames[animation.CurrentFrameIndex]
	}

	return rect
}

func (animation SpriteAnimation) GetCurrentSprite() render.Sprite {
	return render.Sprite{
		Texture: animation.Texture,
		Rect:    animation.GetCurrentRect(),
	}
}

func (animation SpriteAnimation) Copy() IAnimation {
	return &SpriteAnimation{
		Animation: Animation{
			State: AnimationRunning,
		},

		Texture:       animation.Texture,
		Frames:        animation.Frames,
		FrameDuration: animation.FrameDuration,
	}
}

func (animation SpriteAnimation) Reverse() IAnimation {
	copy := &SpriteAnimation{
		Animation: Animation{
			State: AnimationRunning,
		},

		Texture:       animation.Texture,
		FrameDuration: animation.FrameDuration,
	}

	utils.ForEachBackwards(animation.Frames, func(frame math.Box) {
		copy.Frames = append(copy.Frames, frame)
	})

	return copy
}
