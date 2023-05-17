package render

import (
	"github.com/nosimplegames/ns-framework/math"
)

type Slice9Factory struct {
	Texture Texture
	Top     int
	Right   int
	Bottom  int
	Left    int
}

func (factory Slice9Factory) Create() Slice9 {
	slice9 := Slice9{}

	slice9.Textures = factory.GetTextures()

	return slice9
}

func (factory Slice9Factory) GetAreas() [9]math.Box {
	size := GetTextureSize(factory.Texture)

	return Slice9AreasFactory{
		Left:   float64(factory.Left),
		Top:    float64(factory.Top),
		Right:  float64(factory.Right),
		Bottom: float64(factory.Bottom),
		Width:  size.X,
		Height: size.Y,
	}.Create()
}

func (factory Slice9Factory) GetTextures() [9]Texture {
	areas := factory.GetAreas()
	textures := [9]Texture{}

	for areaIndex, area := range areas {
		texture := factory.Texture.SubImage(area.ImageRect()).(Texture)
		textures[areaIndex] = texture
	}

	return textures
}
