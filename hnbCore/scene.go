package hnbCore

import "github.com/nosimplegames/ns-framework/hnbPhysics"

type Scene struct {
	Entity
}

func (scene Scene) RegisterCollisionables() {
	registerEntityCollisionables(&scene)
}

func registerEntityCollisionables(entity IEntity) {
	for _, child := range entity.GetChildren() {
		childAsIEntity := child.(IEntity)
		hnbPhysics.AddCollisionable(childAsIEntity)
		registerEntityCollisionables(childAsIEntity)
	}
}
