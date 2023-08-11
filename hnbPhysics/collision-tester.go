package hnbPhysics

import "github.com/nosimplegames/ns-framework/hnbMath"

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
		Another: tester.RightCollisionable,
	}

	rightCollision := Collision{
		AnotherCollisionMask: tester.LeftCollisionable.GetCollisionMask(),
		CollisionResolverCalculator: CollisionResolverCalculator{
			LeftAABB:  rightAABB,
			RightAABB: leftAABB,
		},
		Another: tester.LeftCollisionable,
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

func (tester CollisionTester) GetLeftAABB() hnbMath.Box {
	return hnbMath.Box{
		Size:     tester.LeftCollisionable.GetSize(),
		Position: tester.LeftCollisionable.GetPosition(),
	}
}

func (tester CollisionTester) GetRightAABB() hnbMath.Box {
	return hnbMath.Box{
		Size:     tester.RightCollisionable.GetSize(),
		Position: tester.RightCollisionable.GetPosition(),
	}
}
