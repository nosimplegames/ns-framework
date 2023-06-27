package oldanimations

import "github.com/nosimplegames/ns-framework/utils"

type EmptyAnimation struct {
	Animation

	Duration    float64
	CurrentTime float64
}

func (animation *EmptyAnimation) Update() {
	if !animation.IsRunning() {
		return
	}

	animation.CurrentTime += utils.FrameTime
	hasFinished := animation.CurrentTime >= animation.Duration

	if hasFinished {
		animation.Stop()
	}
}

func (animation EmptyAnimation) Copy() IAnimation {
	return &EmptyAnimation{
		Duration: animation.Duration,
	}
}

func (animation EmptyAnimation) Reverse() IAnimation {
	return animation.Copy()
}
