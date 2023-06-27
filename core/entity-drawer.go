package core

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type EntityDrawer struct {
	Entity    IEntity
	Transform math.Transform
	Target    render.RenderTarget
}

func (drawer EntityDrawer) Draw() {
	canDraw := drawer.Entity.IsAlive() && drawer.Entity.IsVisible()
	if !canDraw {
		return
	}

	transform := drawer.Entity.GetTransform()
	transform.Concat(drawer.Transform)

	mustDrawEntityBeforeChildren := drawer.Entity.GetDrawPolicy() == DrawBeforeChildren

	if mustDrawEntityBeforeChildren {
		drawer.Entity.Draw(drawer.Target, transform)
	}

	for _, child := range drawer.Entity.GetChildren() {
		EntityDrawer{
			Entity:    child.(IEntity),
			Transform: transform,
			Target:    drawer.Target,
		}.Draw()
	}

	if !mustDrawEntityBeforeChildren {
		drawer.Entity.Draw(drawer.Target, transform)
	}
}
