package oldanimations

import "github.com/nosimplegames/ns-framework/utils"

type AnimationList struct {
	Animation

	Animations            []IAnimation
	CurrentAnimationIndex int
	CurrentAnimation      IAnimation
}

func (list *AnimationList) AddAnimation(animation IAnimation) {
	list.Animations = append(list.Animations, animation)

	mustInitCurrentAnimation := list.CurrentAnimation == nil

	if mustInitCurrentAnimation {
		list.CurrentAnimation = animation
	}
}

func (list *AnimationList) Update() {
	if !list.IsRunning() {
		return
	}

	list.CurrentAnimation.Update()
	list.ChangeCurrentAnimationIfRequired()
}

func (list *AnimationList) ChangeCurrentAnimationIfRequired() {
	needToChangeCurrentAnimation := !list.CurrentAnimation.IsAlive()

	if needToChangeCurrentAnimation {
		list.CurrentAnimationIndex++

		needToStopAnimation := list.CurrentAnimationIndex >= len(list.Animations)

		if needToStopAnimation {
			list.Stop()
		} else {
			list.CurrentAnimation = list.Animations[list.CurrentAnimationIndex]
		}
	}
}

func (list AnimationList) Copy() IAnimation {
	copy := &AnimationList{}

	copy.State = AnimationRunning

	for _, animation := range list.Animations {
		copy.AddAnimation(animation.Copy())
	}

	return copy
}

func (list AnimationList) Reverse() IAnimation {
	copy := &AnimationList{}

	copy.State = AnimationRunning

	utils.ForEachBackwards(list.Animations, func(animation IAnimation) {
		copy.AddAnimation(animation.Copy())
	})

	return copy
}
