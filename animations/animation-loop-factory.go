package animations

type AnimationLoopFactory struct {
	Animation             IAnimation
	LoopCount             int
	IncludeBackwards      bool
	TimeBetweenAnimations float64
}

func (factory AnimationLoopFactory) Create() AnimationList {
	animationList := AnimationList{}
	animationList.State = AnimationRunning
	emptyAnimation := EmptyAnimation{
		Duration: factory.TimeBetweenAnimations,
	}

	for i := 0; i < factory.LoopCount; i++ {
		animationList.AddAnimation(factory.Animation.Copy())

		isLastAnimation := i == factory.LoopCount-1 && !factory.IncludeBackwards
		mustAddEmptyAnimation := factory.TimeBetweenAnimations != 0 && !isLastAnimation

		if mustAddEmptyAnimation {
			animationList.AddAnimation(emptyAnimation.Copy())
		}

		if factory.IncludeBackwards {
			animationList.AddAnimation(factory.Animation.Reverse())

			isLastBackwardsAnimation := i == factory.LoopCount-1
			mustAddEmptyAnimation := factory.TimeBetweenAnimations != 0 && !isLastBackwardsAnimation

			if mustAddEmptyAnimation {
				animationList.AddAnimation(emptyAnimation.Copy())
			}
		}
	}

	return animationList
}
