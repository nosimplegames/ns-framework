package hnbPhysics

import "github.com/nosimplegames/ns-framework/hnbMath"

type CollisionableFactory struct {
	Size              hnbMath.Vector
	CanCollide        bool
	CollisionMask     string
	CollisioningMasks []string
}

func (factory CollisionableFactory) Init(collisionable *Collisionable) {
	collisionable.size = factory.Size
	collisionable.canCollide = factory.CanCollide
	collisionable.collisionMask = factory.CollisionMask
	collisionable.collisioningMasks = factory.CollisioningMasks
}

func (factory CollisionableFactory) Create() *Collisionable {
	collisionable := &Collisionable{}

	factory.Init(collisionable)

	return collisionable
}
