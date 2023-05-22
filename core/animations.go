package core

import "github.com/nosimplegames/ns-framework/utils"

type Animations struct {
	animations []IAnimation
}

func (animations *Animations) AddAnimation(animation IAnimation) {
	animations.animations = append(animations.animations, animation)
}

func (animations *Animations) Update() {
	for _, animation := range animations.animations {
		if !animation.IsAlive() {
			continue
		}

		animation.Apply()
		animation.Update()

		if animation.HasFinished() {
			if animation.WillReplay() {
				animation.Restart()
			} else {
				animation.Stop()
			}
		}
	}

	utils.RemoveDead(&animations.animations)
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
