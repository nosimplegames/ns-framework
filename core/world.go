package core

import (
	"github.com/nosimplegames/ns-framework/physics"
	"github.com/nosimplegames/ns-framework/utils"
)

type World struct {
	Collisionables []physics.ICollisionable
}

func (world *World) AddCollisinable(collisionable physics.ICollisionable) {
	world.Collisionables = append(world.Collisionables, collisionable)
}

func (world *World) Update() {
	world.TestCollisions()
	world.RemoveDead()
}

func (world *World) TestCollisions() {
	for i := 0; i < len(world.Collisionables); i++ {
		leftCollisionable := world.Collisionables[i]

		mayTestCollisions := leftCollisionable.IsAlive() && leftCollisionable.CanCollide()

		if !mayTestCollisions {
			continue
		}

		for j := i + 1; j < len(world.Collisionables); j++ {
			rightCollisionable := world.Collisionables[j]
			mustTestCollision := rightCollisionable.IsAlive() &&
				rightCollisionable.CanCollide() &&
				leftCollisionable.CanCollideWith(rightCollisionable.GetCollisionMask())

			if mustTestCollision {
				collision := physics.Collision{
					LeftCollisionable:  leftCollisionable,
					RightCollisionable: rightCollisionable,
				}

				collision.ReportIfHapenning()
			}
		}
	}
}

func (world *World) RemoveDead() {
	utils.RemoveDead(&world.Collisionables)
}

var world *World = nil

func GetWorld() *World {
	needToInitWorld := world == nil

	if needToInitWorld {
		world = &World{}
	}

	return world
}
