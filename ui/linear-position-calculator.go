package ui

import "github.com/nosimplegames/ns-framework/utils"

type JustifyContentPolicy = string

const (
	JustifyContentStart        JustifyContentPolicy = ""
	JustifyContentSpacebetween                      = "Spacebetween"
	JustifyContentCenter                            = "Center"
)

type LinearPositionCalculator struct {
	Size                 float64
	JustifyContentPolicy JustifyContentPolicy
	ElementsSizes        []float64
	Gap                  float64
}

func (calculator LinearPositionCalculator) Calculate() []float64 {
	positions := []float64{}
	spaceBetweenElements := calculator.CalculateSpaceBetweenElements()
	offsetPosition := calculator.CalculateOffsetPosition()

	for _, elementSize := range calculator.ElementsSizes {
		elementPosition := offsetPosition + elementSize*0.5
		positions = append(positions, elementPosition)
		offsetPosition += elementSize + spaceBetweenElements
	}

	return positions
}

func (calculator LinearPositionCalculator) CalculateElementsSize() float64 {
	return utils.Reduce(
		calculator.ElementsSizes,
		func(accumulator float64, elementSize float64) float64 {
			return accumulator + elementSize
		},
		0.0,
	)
}

func (calculator LinearPositionCalculator) CalculateElementsSizeWithGap() float64 {
	elementsSizeWithGap := utils.Reduce(
		calculator.ElementsSizes,
		func(accumulator float64, elementSize float64) float64 {
			return accumulator + elementSize + calculator.Gap
		},
		0.0,
	)

	elementsSizeWithGap -= calculator.Gap
	return elementsSizeWithGap
}

func (calculator LinearPositionCalculator) CalculateFreeSpace() float64 {
	elementsSize := calculator.CalculateElementsSize()
	freeSpace := calculator.Size - elementsSize

	return freeSpace
}

func (calculator LinearPositionCalculator) CalculateSpaceBetweenElements() float64 {
	mayUtilizeGap := calculator.JustifyContentPolicy != JustifyContentSpacebetween

	if mayUtilizeGap {
		return calculator.Gap
	}

	freeSpace := calculator.CalculateFreeSpace()
	spaceBetweenElements := freeSpace / float64(len(calculator.ElementsSizes)-1)

	return spaceBetweenElements
}

func (calculator LinearPositionCalculator) CalculateOffsetPosition() float64 {
	isContentCentered := calculator.JustifyContentPolicy == JustifyContentCenter

	if isContentCentered {
		elementsSize := calculator.CalculateElementsSizeWithGap()
		return -elementsSize * 0.5
	}

	return -calculator.Size * 0.5
}
