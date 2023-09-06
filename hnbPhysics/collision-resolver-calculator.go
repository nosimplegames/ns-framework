package hnbPhysics

import (
	"math"

	"github.com/nosimplegames/ns-framework/hnbMath"
)

type CollisionResolverCalculator struct {
	LeftAABB  hnbMath.Box
	RightAABB hnbMath.Box
}

func (calculator CollisionResolverCalculator) CalculateXResolution() float64 {
	lineA := hnbMath.LinearLine{
		Start:  calculator.LeftAABB.Position.X - calculator.LeftAABB.Size.X*0.5,
		Length: calculator.LeftAABB.Size.X,
	}
	lineB := hnbMath.LinearLine{
		Start:  calculator.RightAABB.Position.X - calculator.RightAABB.Size.X*0.5,
		Length: calculator.RightAABB.Size.X,
	}
	penetration := hnbMath.LinearPenetrationCalculator{
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
	lineA := hnbMath.LinearLine{
		Start:  calculator.LeftAABB.Position.Y - calculator.LeftAABB.Size.Y*0.5,
		Length: calculator.LeftAABB.Size.Y,
	}
	lineB := hnbMath.LinearLine{
		Start:  calculator.RightAABB.Position.Y - calculator.RightAABB.Size.Y*0.5,
		Length: calculator.RightAABB.Size.Y,
	}
	penetration := hnbMath.LinearPenetrationCalculator{
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
	lineA hnbMath.LinearLine,
	lineB hnbMath.LinearLine,
) bool {
	return lineB.Start <= lineA.Start && lineB.End() <= lineA.End()
}

func (calculator CollisionResolverCalculator) GetMinVectorResolver() hnbMath.Vector {
	xResolution := calculator.CalculateXResolution()
	yResolution := calculator.CalculateYResolution()

	isXLowerThanY := math.Abs(xResolution) < math.Abs(yResolution)

	if isXLowerThanY {
		return hnbMath.Vector{
			X: xResolution,
		}
	}

	return hnbMath.Vector{
		Y: yResolution,
	}
}
