package hnbPhysics

type Collision struct {
	AnotherCollisionMask        string
	CollisionResolverCalculator CollisionResolverCalculator
	Another                     ICollisionable
}
