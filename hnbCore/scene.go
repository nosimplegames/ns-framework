package hnbCore

import "github.com/nosimplegames/ns-framework/hnbPhysics"

type Scene struct {
	Entity
}

func (scene Scene) RegisterCollisionables() {
	registerEntityCollisionables(&scene)
}

func registerEntityCollisionables(entity IEntity) {
	for _, childNode := range entity.GetChildren() {
		child := childNode.(IEntity)

		if child.CanCollide() {
			hnbPhysics.AddCollisionable(child)
		}

		registerEntityCollisionables(child)
	}
}
