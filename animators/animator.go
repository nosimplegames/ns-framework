package animators

import "github.com/nosimplegames/ns-framework/animations"

type AnimateTargetFunc[T any, A animations.IAnimation] func(target T, animation A)

type Animator[T any, A animations.IAnimation] struct {
	Animation     A
	Target        T
	AnimateTarget AnimateTargetFunc[T, A]
}

func (animator Animator[T, A]) IsAlive() bool {
	return animator.Animation.IsAlive()
}

func (animator Animator[T, A]) IsRunning() bool {
	return animator.Animation.IsRunning()
}

func (animator *Animator[T, A]) Update() {
	animator.Animation.Update()

	if animator.Animation.IsRunning() {
		animator.AnimateTarget(animator.Target, animator.Animation)
	}
}

func (animator Animator[T, A]) Copy() animations.IAnimation {
	return &Animator[T, A]{
		Animation:     animator.Animation.Copy().(A),
		Target:        animator.Target,
		AnimateTarget: animator.AnimateTarget,
	}
}

func (animator Animator[T, A]) Reverse() animations.IAnimation {
	return &Animator[T, A]{
		Animation:     animator.Animation.Reverse().(A),
		Target:        animator.Target,
		AnimateTarget: animator.AnimateTarget,
	}
}
