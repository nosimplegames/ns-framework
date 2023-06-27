package ui

type Padding struct {
	Top    float64
	Right  float64
	Bottom float64
	Left   float64
}

func SamePadding(padding float64) Padding {
	return Padding{
		Top:    padding,
		Right:  padding,
		Bottom: padding,
		Left:   padding,
	}
}

func HorizontalVerticalPadding(horizontal float64, vertical float64) Padding {
	return Padding{
		Top:    vertical,
		Right:  horizontal,
		Bottom: vertical,
		Left:   horizontal,
	}
}
