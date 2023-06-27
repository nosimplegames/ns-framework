package oldanimations

import "github.com/nosimplegames/ns-framework/utils"

type NumberAnimationFactory struct {
	InitialValue float64
	TargetValue  float64
	Duration     float64
}

func (factory NumberAnimationFactory) Create() NumberAnimation {
	animation := NumberAnimation{}

	animation.CurrentValue = factory.InitialValue
	animation.TargetValue = factory.TargetValue
	animation.FrameStep = factory.GetFrameStep()

	return animation
}

func (factory NumberAnimationFactory) GetFrameStep() float64 {
	deltaValue := factory.TargetValue - factory.InitialValue
	frameStep := utils.FrameTime * deltaValue / factory.Duration

	return frameStep
}
