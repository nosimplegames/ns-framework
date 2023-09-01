package hnbCore

import "github.com/nosimplegames/ns-framework/hnbUtils"

type Animations struct {
	animations []IAnimation
}

func (animations *Animations) AddAnimation(animation IAnimation) {
	animations.animations = append(animations.animations, animation)
}

func (animations *Animations) UpdateFrame() {
	for _, animation := range animations.animations {
		if !animation.IsAlive() {
			continue
		}

		if animation.IsSleeping() {
			animations.handleSleepingAnimation(animation)
		} else if animation.IsRunning() {
			animations.handleRunningAnimation(animation)
		}
	}

	hnbUtils.RemoveDead(&animations.animations)
}

func (animations Animations) handleSleepingAnimation(animation IAnimation) {
	animation.KeepSleeping()

	if animation.MustWakeUp() {
		animation.Restart()
		animations.handleRunningAnimation(animation)
	}
}

func (animations Animations) handleRunningAnimation(animation IAnimation) {
	animation.Apply()
	animation.UpdateFrame()

	if animation.HasFinished() {
		if animation.WillReplay() {
			if animation.WillSleep() {
				animation.StartSleeping()
			} else {
				animation.Restart()
			}
		} else {
			animation.Stop()
		}
	}
}

var animations *Animations = nil

func GetAnimations() *Animations {
	needToInitAnimations := animations == nil

	if needToInitAnimations {
		animations = &Animations{}
	}

	return animations
}

func AddAnimation(animation IAnimation) {
	GetAnimations().AddAnimation(animation)
}
