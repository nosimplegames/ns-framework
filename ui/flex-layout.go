package ui

import "github.com/nosimplegames/ns-framework/math"

type FlexLayout struct {
	LayoutDirection      LayoutDirection
	JustifyContentPolicy JustifyContentPolicy
	AlignItemsPolicy     JustifyContentPolicy

	Size math.Vector
	Gap  math.Vector
}

const (
	FlexRow    LayoutDirection = ""
	FlexColumn                 = "FlexColumn"
)

func (layout *FlexLayout) SetSize(size math.Vector) {
	layout.Size = size
}

func (layout FlexLayout) GetPositions(elements []LayoutElement) []math.Vector {
	positions := []math.Vector{}
	xPositions := layout.GetXPositions(elements)
	yPositions := layout.GetYPositions(elements)

	for index, xPosition := range xPositions {
		position := math.Vector{
			X: xPosition,
			Y: yPositions[index],
		}
		positions = append(positions, position)
	}

	return positions
}

func (layout FlexLayout) GetXPositions(elements []LayoutElement) []float64 {
	elementSizes := []float64{}

	for _, element := range elements {
		elementSizes = append(elementSizes, element.Size.X)
	}

	justifyContentPolicy := layout.GetXContentPolicy()

	positions := []float64{}

	if layout.LayoutDirection == FlexRow {
		positions = layout.GetTPositions(
			layout.Size.X,
			layout.Gap.X,
			elementSizes,
			justifyContentPolicy,
		)
	} else {
		positions = layout.GetTIndividualPositions(
			layout.Size.X,
			layout.Gap.X,
			elementSizes,
			justifyContentPolicy,
		)
	}

	return positions
}

func (layout FlexLayout) GetXSizes(elements []LayoutElement) []float64 {
	xSizes := []float64{}

	for _, element := range elements {
		xSizes = append(xSizes, element.Size.X)
	}

	return xSizes
}

func (layout FlexLayout) GetYSizes(elements []LayoutElement) []float64 {
	ySizes := []float64{}

	for _, element := range elements {
		ySizes = append(ySizes, element.Size.Y)
	}

	return ySizes
}

func (layout FlexLayout) GetTPositions(
	size float64,
	gap float64,
	elementsSizes []float64,
	justifyContentPolicy JustifyContentPolicy,
) []float64 {
	linearPositionCalculator := LinearPositionCalculator{
		Size:                 size,
		JustifyContentPolicy: justifyContentPolicy,
		ElementsSizes:        elementsSizes,
		Gap:                  gap,
	}

	positions := linearPositionCalculator.Calculate()

	return positions
}

func (layout FlexLayout) GetTIndividualPositions(
	size float64,
	gap float64,
	elementsSizes []float64,
	justifyContentPolicy JustifyContentPolicy,
) []float64 {
	positions := []float64{}

	for _, elementSize := range elementsSizes {
		linearPositionCalculator := LinearPositionCalculator{
			Size:                 size,
			JustifyContentPolicy: justifyContentPolicy,
			ElementsSizes:        []float64{elementSize},
			Gap:                  gap,
		}

		elementPositions := linearPositionCalculator.Calculate()
		positions = append(positions, elementPositions[0])
	}

	return positions
}

func (layout FlexLayout) GetXContentPolicy() JustifyContentPolicy {
	shouldUseJustifyContentPolicy := layout.LayoutDirection == FlexRow

	if shouldUseJustifyContentPolicy {
		return layout.JustifyContentPolicy
	}

	return layout.AlignItemsPolicy
}

func (layout FlexLayout) GetYContentPolicy() JustifyContentPolicy {
	shouldUseJustifyContentPolicy := layout.LayoutDirection == FlexColumn

	if shouldUseJustifyContentPolicy {
		return layout.JustifyContentPolicy
	}

	return layout.AlignItemsPolicy
}

func (layout FlexLayout) GetYPositions(elements []LayoutElement) []float64 {
	elementSizes := []float64{}

	for _, element := range elements {
		elementSizes = append(elementSizes, element.Size.Y)
	}

	justifyContentPolicy := layout.GetYContentPolicy()

	positions := []float64{}

	if layout.LayoutDirection == FlexColumn {
		positions = layout.GetTPositions(
			layout.Size.Y,
			layout.Gap.Y,
			elementSizes,
			justifyContentPolicy,
		)
	} else {
		positions = layout.GetTIndividualPositions(
			layout.Size.Y,
			layout.Gap.Y,
			elementSizes,
			justifyContentPolicy,
		)
	}

	return positions
}

func (layout FlexLayout) GetXSizePolicy() SizeCalculationPolicy {
	policy := AllElementsSizeCalculationPolicy

	shouldUseBiggerElement := layout.LayoutDirection == FlexColumn

	if shouldUseBiggerElement {
		policy = BiggerElementSizeCalculationPolicy
	}

	return policy
}

func (layout FlexLayout) GetYSizePolicy() SizeCalculationPolicy {
	policy := AllElementsSizeCalculationPolicy

	shouldUseBiggerElement := layout.LayoutDirection == FlexRow

	if shouldUseBiggerElement {
		policy = BiggerElementSizeCalculationPolicy
	}

	return policy
}

func (layout FlexLayout) GetContentSize(elements []LayoutElement) math.Vector {
	size := math.Vector{}

	size.X = LinearSizeCalculator{
		CalculationPolicy: layout.GetXSizePolicy(),
		ElementsSizes:     layout.GetXSizes(elements),
		Gap:               layout.Gap.X,
	}.Calculate()

	size.Y = LinearSizeCalculator{
		CalculationPolicy: layout.GetYSizePolicy(),
		ElementsSizes:     layout.GetYSizes(elements),
		Gap:               layout.Gap.Y,
	}.Calculate()

	return size
}
