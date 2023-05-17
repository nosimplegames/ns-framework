package physics

import "github.com/nosimplegames/ns-framework/math"

type CollisionResolverCalculator struct {
	LeftAABB  math.Box
	RightAABB math.Box
}

func (calculator CollisionResolverCalculator) CalculateXResolution() float64 {
	lineA := math.LinearLine{
		Start:  calculator.LeftAABB.Position.X - calculator.LeftAABB.Size.X*0.5,
		Length: calculator.LeftAABB.Size.X,
	}
	lineB := math.LinearLine{
		Start:  calculator.RightAABB.Position.X - calculator.RightAABB.Size.X*0.5,
		Length: calculator.RightAABB.Size.X,
	}
	penetration := math.LinearPenetrationCalculator{
		LineA: lineA,
		LineB: lineB,
	}.Calculate()

	isPositive := calculator.IsResolutionPositive(lineA, lineB)

	if !isPositive {
		return -penetration
	}

	return penetration
}

func (calculator CollisionResolverCalculator) CalculateYResolution() float64 {
	lineA := math.LinearLine{
		Start:  calculator.LeftAABB.Position.Y - calculator.LeftAABB.Size.Y*0.5,
		Length: calculator.LeftAABB.Size.Y,
	}
	lineB := math.LinearLine{
		Start:  calculator.RightAABB.Position.Y - calculator.RightAABB.Size.Y*0.5,
		Length: calculator.RightAABB.Size.Y,
	}
	penetration := math.LinearPenetrationCalculator{
		LineA: lineA,
		LineB: lineB,
	}.Calculate()

	isPositive := calculator.IsResolutionPositive(lineA, lineB)

	if !isPositive {
		return -penetration
	}

	return penetration
}

func (calculator CollisionResolverCalculator) IsResolutionPositive(
	lineA math.LinearLine,
	lineB math.LinearLine,
) bool {
	return lineB.Start <= lineA.Start && lineB.End() <= lineA.End()
}
