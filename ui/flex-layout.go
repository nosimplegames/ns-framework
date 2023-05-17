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

	positions := layout.GetTPositions(
		layout.Size.X,
		layout.Gap.X,
		elementSizes,
		justifyContentPolicy,
	)

	return positions
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

	positions := layout.GetTPositions(
		layout.Size.Y,
		layout.Gap.Y,
		elementSizes,
		justifyContentPolicy,
	)

	return positions
}
