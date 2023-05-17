package physics

import (
	"github.com/nosimplegames/ns-framework/math"
)

type CollisionTester struct {
	LeftCollisionable  ICollisionable
	RightCollisionable ICollisionable
}

func (tester CollisionTester) IsHappening() bool {
	leftAABB := tester.GetLeftAABB()
	rightAABB := tester.GetRightAABB()

	return leftAABB.IsColliding(rightAABB)
}

func (tester CollisionTester) Report() {
	leftAABB := tester.GetLeftAABB()
	rightAABB := tester.GetRightAABB()

	leftCollision := Collision{
		AnotherCollisionMask: tester.RightCollisionable.GetCollisionMask(),
		CollisionResolverCalculator: CollisionResolverCalculator{
			LeftAABB:  leftAABB,
			RightAABB: rightAABB,
		},
	}

	rightCollision := Collision{
		AnotherCollisionMask: tester.RightCollisionable.GetCollisionMask(),
		CollisionResolverCalculator: CollisionResolverCalculator{
			LeftAABB:  rightAABB,
			RightAABB: leftAABB,
		},
	}

	tester.LeftCollisionable.OnCollision(leftCollision)
	tester.RightCollisionable.OnCollision(rightCollision)
}

func (tester CollisionTester) ReportIfHapenning() {
	if !tester.IsHappening() {
		return
	}

	tester.Report()
}

func (tester CollisionTester) GetLeftAABB() math.Box {
	return math.Box{
		Size:     tester.LeftCollisionable.GetSize(),
		Position: tester.LeftCollisionable.GetPosition(),
	}
}

func (tester CollisionTester) GetRightAABB() math.Box {
	return math.Box{
		Size:     tester.RightCollisionable.GetSize(),
		Position: tester.RightCollisionable.GetPosition(),
	}
}
