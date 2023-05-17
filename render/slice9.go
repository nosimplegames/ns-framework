package render

import (
	"github.com/nosimplegames/ns-framework/math"
)

type Slice9 struct {
	Textures [9]Texture
}

func (slice9 Slice9) Compose(targetSize math.Vector) Texture {
	texture := TextureFactory{
		TargetSize: targetSize,
	}.Create()

	areas := slice9.GetTargetAreas(targetSize)

	for areaIndex, area := range areas {
		TextureRepeater{
			Target:  texture,
			Texture: slice9.Textures[areaIndex],
			Area:    area,
		}.Repeat()
	}

	return texture
}

func (slice9 Slice9) GetTargetAreas(targetSize math.Vector) [9]math.Box {
	topLeftCornerTextureSize := GetTextureSize(slice9.Textures[0])
	bottomRightCornerTextureSize := GetTextureSize(slice9.Textures[8])

	return Slice9AreasFactory{
		Left:   topLeftCornerTextureSize.X,
		Top:    topLeftCornerTextureSize.Y,
		Right:  targetSize.X - bottomRightCornerTextureSize.X,
		Bottom: targetSize.Y - bottomRightCornerTextureSize.Y,
		Width:  targetSize.X,
		Height: targetSize.Y,
	}.Create()
}
