package math

import "math"

type LinearPenetrationCalculator struct {
	LineA LinearLine
	LineB LinearLine
}

func (calculator LinearPenetrationCalculator) Calculate() float64 {
	penetration := math.Abs(calculator.LineA.End() - calculator.LineB.Start)

	return penetration
}
