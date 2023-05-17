package animators

import "github.com/nosimplegames/ns-framework/animations"

type AlphaAnimatorFactory struct {
	Animation animations.NumberAnimation
}

func (factory AlphaAnimatorFactory) Create(target IColorable) *AlphaAnimator {
	animator := &AlphaAnimator{}

	alphaAnimation := factory.Animation.Copy().(*animations.NumberAnimation)
	animator.Animation = alphaAnimation

	animator.Target = target

	animator.AnimateTarget = func(target IColorable, animation *animations.NumberAnimation) {
		target.SetAlpha(animation.CurrentValue)
	}

	return animator
}
