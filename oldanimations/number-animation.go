package oldanimations

type NumberAnimation struct {
	Animation

	InitialValue float64
	TargetValue  float64
	FrameStep    float64
	CurrentValue float64
}

func (animation *NumberAnimation) Update() {
	animation.CurrentValue += animation.FrameStep

	if animation.HasFinished() {
		animation.CurrentValue = animation.TargetValue
		animation.Stop()
	}
}

func (animation NumberAnimation) HasFinished() bool {
	targetGreaterThanInitial := animation.TargetValue > animation.InitialValue

	if targetGreaterThanInitial {
		return animation.CurrentValue >= animation.TargetValue
	}

	return animation.CurrentValue <= animation.TargetValue
}

func (animation NumberAnimation) Copy() IAnimation {
	copy := &NumberAnimation{
		InitialValue: animation.InitialValue,
		TargetValue:  animation.TargetValue,
		CurrentValue: animation.InitialValue,
		FrameStep:    animation.FrameStep,
	}

	copy.State = AnimationRunning

	return copy
}

func (animation NumberAnimation) Reverse() IAnimation {
	copy := &NumberAnimation{
		InitialValue: animation.TargetValue,
		TargetValue:  animation.InitialValue,
		CurrentValue: animation.TargetValue,
		FrameStep:    -animation.FrameStep,
	}

	copy.State = AnimationRunning

	return copy
}
