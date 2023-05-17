package math

type LinearLine struct {
	Start  float64
	Length float64
}

func (line LinearLine) End() float64 {
	return line.Start + line.Length
}
