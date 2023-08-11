package hnbRender

import "github.com/nosimplegames/ns-framework/hnbMath"

type Slice9AreasFactory struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
	Width  float64
	Height float64
}

func (factory Slice9AreasFactory) Create() [9]hnbMath.Box {
	return [9]hnbMath.Box{
		hnbMath.BoxFromCorners(0, 0, factory.Left, factory.Top),
		hnbMath.BoxFromCorners(factory.Left, 0, factory.Right, factory.Top),
		hnbMath.BoxFromCorners(factory.Right, 0, factory.Width, factory.Top),

		hnbMath.BoxFromCorners(0, factory.Top, factory.Left, factory.Bottom),
		hnbMath.BoxFromCorners(factory.Left, factory.Top, factory.Right, factory.Bottom),
		hnbMath.BoxFromCorners(factory.Right, factory.Top, factory.Width, factory.Bottom),

		hnbMath.BoxFromCorners(0, factory.Bottom, factory.Left, factory.Height),
		hnbMath.BoxFromCorners(factory.Left, factory.Bottom, factory.Right, factory.Height),
		hnbMath.BoxFromCorners(factory.Right, factory.Bottom, factory.Width, factory.Height),
	}
}
