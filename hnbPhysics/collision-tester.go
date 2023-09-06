package hnbPhysics

type CollisionTester struct {
	LeftCollisionable  ICollisionable
	RightCollisionable ICollisionable
}

func (tester CollisionTester) IsHappening() bool {
	leftAABB := tester.LeftCollisionable.GetAABB()
	rightAABB := tester.RightCollisionable.GetAABB()

	return leftAABB.IsColliding(rightAABB)
}

func (tester CollisionTester) Report() {
	leftAABB := tester.LeftCollisionable.GetAABB()
	rightAABB := tester.RightCollisionable.GetAABB()

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
