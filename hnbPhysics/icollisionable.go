package hnbPhysics

import "github.com/nosimplegames/ns-framework/hnbMath"

type ICollisionable interface {
	hnbMath.ITransformable

	GetSize() hnbMath.Vector
	GetPosition() hnbMath.Vector
	GetCollisionMask() string
	CanCollide() bool
	CanCollideWith(collisionMask string) bool
	OnCollision(collision Collision)
	IsAlive() bool
}
