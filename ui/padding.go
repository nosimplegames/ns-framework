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
