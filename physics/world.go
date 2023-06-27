package physics

import (
	"github.com/nosimplegames/ns-framework/render"
	"github.com/nosimplegames/ns-framework/utils"
)

type World struct {
	Collisionables []ICollisionable
}

func (world *World) AddCollisinable(collisionable ICollisionable) {
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
			canCollisionablesCollide := leftCollisionable.CanCollideWith(rightCollisionable.GetCollisionMask()) ||
				rightCollisionable.CanCollideWith(leftCollisionable.GetCollisionMask())
			canRightCollide := rightCollisionable.IsAlive() &&
				rightCollisionable.CanCollide()
			mustTestCollision := canRightCollide && canCollisionablesCollide

			if mustTestCollision {
				tester := CollisionTester{
					LeftCollisionable:  leftCollisionable,
					RightCollisionable: rightCollisionable,
				}

				tester.ReportIfHapenning()
			}
		}
	}
}

func (world *World) RemoveDead() {
	utils.RemoveDead(&world.Collisionables)
}

func (world World) Draw(target render.RenderTarget) {
	for _, collisionable := range world.Collisionables {
		render.Box{
			Position: collisionable.GetPosition(),
			Size:     collisionable.GetSize(),
			Target:   target,
		}.Render()
	}
}

var world *World = nil

func GetWorld() *World {
	needToInitWorld := world == nil

	if needToInitWorld {
		world = &World{}
	}

	return world
}

func AddCollisionable(collisionable ICollisionable) {
	world := GetWorld()
	world.AddCollisinable(collisionable)
}
