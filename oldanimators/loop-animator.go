package oldanimators

import "github.com/nosimplegames/ns-framework/animations"

type LoopAnimator[T any] struct {
	Animator[T, animations.IAnimation]
}

type LoopAnimatorFactory[T any] struct {
	Animation animations.IAnimation
}

func (factory LoopAnimatorFactory[T]) Create(target T) animations.IAnimation {
	animator := &LoopAnimator[T]{}

	animation := factory.Animation.Copy()
	animation.OnStop.AddCallback(func() {
		animator.Target.SetTexture(nil)
		animator.Target.SetRect(nil)
	})
	animator.Animation = spriteAnimation

	animator.Target = target
	animator.Target.SetTexture(spriteAnimation.Texture)
	animator.Target.SetRect(spriteAnimation.GetCurrentRect())

	animator.AnimateTarget = func(target T, animation *animations.SpriteAnimation) {
		target.SetRect(animation.GetCurrentRect())
	}

	return animator
}
