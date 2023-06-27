package oldanimators

import "github.com/nosimplegames/ns-framework/animations"

type IAnimatorFactory[T any] interface {
	Create(target T) animations.IAnimation
}
