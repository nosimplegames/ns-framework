package ui

import "github.com/nosimplegames/ns-framework/math"

type ILayout interface {
	GetPositions(elements []LayoutElement) []math.Vector
	SetSize(math.Vector)
}
