package hnbPhysics

import "github.com/nosimplegames/ns-framework/hnbMath"

type Collisionable struct {
	hnbMath.Transformable

	size              hnbMath.Vector
	canCollide        bool
	collisionMask     string
	collisioningMasks []string
	isDeath           bool
}

func (box *Collisionable) SetSize(size hnbMath.Vector) {
	box.size = size
}

func (box Collisionable) GetSize() hnbMath.Vector {
	return box.size
}

func (collisionable Collisionable) CanCollide() bool {
	return collisionable.canCollide
}

func (collisionable *Collisionable) SetCanCollide(canCollide bool) {
	collisionable.canCollide = canCollide
}

func (collisionable Collisionable) CanCollideWith(another string) bool {
	for _, collisioningMask := range collisionable.collisioningMasks {
		canCollide := another == collisioningMask

		if canCollide {
			return true
		}
	}

	return false
}

func (collisionable *Collisionable) SetCollisioningMasks(masks []string) {
	collisionable.collisioningMasks = masks
}

func (collisionable Collisionable) GetCollisionMask() string {
	return collisionable.collisionMask
}

func (collisionable *Collisionable) SetCollisionMask(mask string) {
	collisionable.collisionMask = mask
}

func (collisionable Collisionable) OnCollision(collision Collision) {
}

func (collisionable Collisionable) IsAlive() bool {
	return !collisionable.isDeath
}

func (collisionable *Collisionable) Die() {
	collisionable.isDeath = true
}

func (collisionable Collisionable) GetAABB() hnbMath.Box {
	return hnbMath.Box{
		Position: collisionable.GetPosition(),
		Size:     collisionable.size,
	}
}

func (collisionable Collisionable) GetPrevAABB() hnbMath.Box {
	return hnbMath.Box{
		Size:     collisionable.size,
		Position: collisionable.GetPrevPosition(),
	}
}
