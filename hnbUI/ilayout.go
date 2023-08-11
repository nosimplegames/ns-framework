package hnbUI

import "github.com/nosimplegames/ns-framework/hnbMath"

type ILayout interface {
	GetPositions(elements []LayoutElement) []hnbMath.Vector
	SetSize(hnbMath.Vector)
	GetContentSize(elements []LayoutElement) hnbMath.Vector
}
