package ui

import (
	"github.com/nosimplegames/ns-framework/core"
	"github.com/nosimplegames/ns-framework/events"
	"github.com/nosimplegames/ns-framework/math"
)

type Component struct {
	core.Entity

	OnSizeChanged events.Signal
}

func (component *Component) SetSize(size math.Vector) {
	currentSize := component.GetSize()

	didSizeChanged := !currentSize.Equals(size)
	if !didSizeChanged {
		return
	}

	component.Entity.SetSize(size)
	component.OnSizeChanged.TFire(size)
}
