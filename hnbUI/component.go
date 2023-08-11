package hnbUI

import (
	"github.com/nosimplegames/ns-framework/hnbCore"
	"github.com/nosimplegames/ns-framework/hnbEvents"
	"github.com/nosimplegames/ns-framework/hnbMath"
)

type Component struct {
	hnbCore.Entity

	OnSizeChanged hnbEvents.Signal
}

func (component *Component) SetSize(size hnbMath.Vector) {
	currentSize := component.GetSize()

	didSizeChanged := !currentSize.Equals(size)
	if !didSizeChanged {
		return
	}

	component.Entity.SetSize(size)
	component.OnSizeChanged.TFire(size)
}
