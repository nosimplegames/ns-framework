package physics

import "github.com/nosimplegames/ns-framework/math"

type ICollisionable interface {
	GetSize() math.Vector
	GetPosition() math.Vector
	GetCollisionMask() string
	CanCollide() bool
	CanCollideWith(collisionMask string) bool
	OnCollision(collision Collision)
	IsAlive() bool
}
