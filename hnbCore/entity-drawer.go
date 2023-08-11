package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type EntityDrawer struct {
	Entity    IEntity
	Transform hnbMath.Transform
	Target    hnbRender.RenderTarget
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
