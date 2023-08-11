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

		animation.Apply()
		animation.UpdateFrame()

		if animation.HasFinished() {
			if animation.WillReplay() {
				animation.Restart()
			} else {
				animation.Stop()
			}
		}
	}

	hnbUtils.RemoveDead(&animations.animations)
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
