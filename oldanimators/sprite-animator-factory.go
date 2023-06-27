package oldanimators

import (
	"github.com/nosimplegames/ns-framework/animations"
)

type SpriteAnimatorFactory struct {
	Animation animations.SpriteAnimation
}

func (factory SpriteAnimatorFactory) Create(target ISprite) animations.IAnimation {
	animator := &SpriteAnimator{}

	spriteAnimation := factory.Animation.Copy().(*animations.SpriteAnimation)
	spriteAnimation.OnStop.AddCallback(func() {
		animator.Target.SetTexture(nil)
		animator.Target.SetRect(nil)
	})
	animator.Animation = spriteAnimation

	animator.Target = target
	animator.Target.SetTexture(spriteAnimation.Texture)
	animator.Target.SetRect(spriteAnimation.GetCurrentRect())

	animator.AnimateTarget = func(target ISprite, animation *animations.SpriteAnimation) {
		target.SetRect(animation.GetCurrentRect())
	}

	return animator
}
