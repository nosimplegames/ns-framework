package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbPhysics"
)

type EntityFactory struct {
	hnbPhysics.CollisionableFactory
}

func (factory EntityFactory) Init(entity *Entity) {
	factory.CollisionableFactory.Init(&entity.Collisionable)
}

func (factory EntityFactory) Create() *Entity {
	entity := &Entity{}

	factory.Init(entity)

	return entity
}
