package render

import "github.com/nosimplegames/ns-framework/math"

type Slice9AreasFactory struct {
	Left   float64
	Top    float64
	Right  float64
	Bottom float64
	Width  float64
	Height float64
}

func (factory Slice9AreasFactory) Create() [9]math.Box {
	return [9]math.Box{
		math.BoxFromCorners(0, 0, factory.Left, factory.Top),
		math.BoxFromCorners(factory.Left, 0, factory.Right, factory.Top),
		math.BoxFromCorners(factory.Right, 0, factory.Width, factory.Top),

		math.BoxFromCorners(0, factory.Top, factory.Left, factory.Bottom),
		math.BoxFromCorners(factory.Left, factory.Top, factory.Right, factory.Bottom),
		math.BoxFromCorners(factory.Right, factory.Top, factory.Width, factory.Bottom),

		math.BoxFromCorners(0, factory.Bottom, factory.Left, factory.Height),
		math.BoxFromCorners(factory.Left, factory.Bottom, factory.Right, factory.Height),
		math.BoxFromCorners(factory.Right, factory.Bottom, factory.Width, factory.Height),
	}
}
