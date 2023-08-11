package hnbMath

import "math"

type LinearPenetrationCalculator struct {
	LineA LinearLine
	LineB LinearLine
}

func (calculator LinearPenetrationCalculator) Calculate() float64 {
	penetrationA := math.Abs(calculator.LineA.End() - calculator.LineB.Start)
	penetrationB := math.Abs(calculator.LineB.End() - calculator.LineA.Start)
	minPenetration := math.Min(penetrationA, penetrationB)

	return minPenetration
}
