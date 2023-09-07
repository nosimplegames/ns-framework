package hnbPhysics

import "github.com/nosimplegames/ns-framework/hnbMath"

type Collision struct {
	AnotherCollisionMask        string
	CollisionResolverCalculator CollisionResolverCalculator
	Another                     ICollisionable
	AABB                        hnbMath.Box
	AnotherAABB                 hnbMath.Box
}
