package physics

import (
	"github.com/nosimplegames/ns-framework/math"
)

type Collision struct {
	LeftCollisionable  ICollisionable
	RightCollisionable ICollisionable
}

func (collision Collision) IsHappening() bool {
	leftAABB := math.Box{
		Size:     collision.LeftCollisionable.GetSize(),
		Position: collision.LeftCollisionable.GetPosition(),
	}

	rightAABB := math.Box{
		Size:     collision.RightCollisionable.GetSize(),
		Position: collision.RightCollisionable.GetPosition(),
	}

	return leftAABB.IsColliding(rightAABB)
}

func (collision Collision) Report() {
	collision.LeftCollisionable.OnCollision(collision.RightCollisionable.GetCollisionMask())
	collision.RightCollisionable.OnCollision(collision.LeftCollisionable.GetCollisionMask())
}

func (collision Collision) ReportIfHapenning() {
	if !collision.IsHappening() {
		return
	}

	collision.Report()
}
