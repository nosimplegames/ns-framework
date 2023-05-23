package physics

type Collision struct {
	AnotherCollisionMask        string
	CollisionResolverCalculator CollisionResolverCalculator
	Another                     ICollisionable
}
