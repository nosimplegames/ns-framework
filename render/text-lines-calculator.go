package render

import (
	"strings"

	"github.com/hajimehoshi/ebiten/text"
	"github.com/nosimplegames/ns-framework/math"
	"golang.org/x/image/font"
)

type TextLinesCalculator struct {
	Text       string
	LineHeight float64
	LineWidth  float64
	FontFace   font.Face
}

func (calculator TextLinesCalculator) Calculate() ([]string, math.Vector) {
	lines := []string{}

	currentLine := ""
	words := strings.Split(calculator.Text, " ")

	for _, word := range words {
		wordToAppend := word
		isFirstWordInLine := len(currentLine) == 0

		if !isFirstWordInLine {
			wordToAppend = " " + word
		}

		bounds := text.BoundString(calculator.FontFace, currentLine+wordToAppend)

		isOutOfBounds := bounds.Max.X > int(calculator.LineWidth)

		if isOutOfBounds {
			lines = append(lines, currentLine)
			currentLine = word
		} else {
			currentLine += wordToAppend
		}
	}

	needToAppendLastLine := currentLine != ""

	if needToAppendLastLine {
		lines = append(lines, currentLine)
	}

	size := math.Vector{
		X: calculator.LineWidth,
		Y: float64(len(lines) * int(calculator.LineHeight)),
	}
	return lines, size
}
