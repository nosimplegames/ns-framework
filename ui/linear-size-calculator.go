package ui

import "math"

type SizeCalculationPolicy int

const (
	BiggerElementSizeCalculationPolicy SizeCalculationPolicy = iota
	AllElementsSizeCalculationPolicy
)

type LinearSizeCalculator struct {
	CalculationPolicy SizeCalculationPolicy
	ElementsSizes     []float64
	Gap               float64
}

func (calculator LinearSizeCalculator) Calculate() float64 {
	size := 0.0

	switch calculator.CalculationPolicy {
	case BiggerElementSizeCalculationPolicy:
		size += calculator.getBiggerElementSize()
	default:
		size += calculator.getElementsSize()
	}

	return size
}

func (calculator LinearSizeCalculator) getBiggerElementSize() float64 {
	biggerElementSize := 0.0

	for _, elementSize := range calculator.ElementsSizes {
		biggerElementSize = math.Max(biggerElementSize, elementSize)
	}

	return biggerElementSize
}

func (calculator LinearSizeCalculator) getElementsSize() float64 {
	size := 0.0

	for _, elementSize := range calculator.ElementsSizes {
		size += elementSize
		size += calculator.Gap
	}

	needToRemoveLastGap := size > 0.0

	if needToRemoveLastGap {
		size -= calculator.Gap
	}

	return size
}
