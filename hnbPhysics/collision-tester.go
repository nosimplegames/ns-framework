package hnbPhysics

import "github.com/nosimplegames/ns-framework/hnbMath"

type CollisionTester struct {
	LeftCollisionable  ICollisionable
	RightCollisionable ICollisionable
}

func (tester CollisionTester) IsHappening() bool {
	leftAABB := hnbMath.Box{
		Position: tester.LeftCollisionable.GetPosition(),
		Size:     tester.LeftCollisionable.GetSize(),
	}
	rightAABB := hnbMath.Box{
		Position: tester.RightCollisionable.GetPosition(),
		Size:     tester.RightCollisionable.GetSize(),
	}

	return leftAABB.IsColliding(rightAABB)
}

func (tester CollisionTester) Report() {
	leftAABB := hnbMath.Box{
		Position: tester.LeftCollisionable.GetPosition(),
		Size:     tester.LeftCollisionable.GetSize(),
	}
	rightAABB := hnbMath.Box{
		Position: tester.RightCollisionable.GetPosition(),
		Size:     tester.RightCollisionable.GetSize(),
	}

	leftCollision := Collision{
		AnotherCollisionMask: tester.RightCollisionable.GetCollisionMask(),
		CollisionResolverCalculator: CollisionResolverCalculator{
			LeftAABB:  leftAABB,
			RightAABB: rightAABB,
		},
		Another:     tester.RightCollisionable,
		AABB:        leftAABB,
		AnotherAABB: rightAABB,
	}

	rightCollision := Collision{
		AnotherCollisionMask: tester.LeftCollisionable.GetCollisionMask(),
		CollisionResolverCalculator: CollisionResolverCalculator{
			LeftAABB:  rightAABB,
			RightAABB: leftAABB,
		},
		Another:     tester.LeftCollisionable,
		AABB:        rightAABB,
		AnotherAABB: leftAABB,
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
