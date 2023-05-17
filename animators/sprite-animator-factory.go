package animators

import "github.com/nosimplegames/ns-framework/animations"

type SpriteAnimatorFactory struct {
	Animation animations.SpriteAnimation
}

func (factory SpriteAnimatorFactory) Create(target ISprite) *SpriteAnimator {
	animator := &SpriteAnimator{}

	spriteAnimation := factory.Animation.Copy().(*animations.SpriteAnimation)
	animator.Animation = spriteAnimation

	animator.Target = target
	animator.Target.SetTexture(spriteAnimation.Texture)

	animator.AnimateTarget = func(target ISprite, animation *animations.SpriteAnimation) {
		target.SetRect(animation.GetCurrentRect())
	}

	return animator
}
