package entities

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type TileMap struct {
	Sprite
}

type TileMapFactory struct {
	TileSet   render.Texture
	TileSize  math.Vector
	TileCount int
	Layer     [][]int
}

func (factory TileMapFactory) Create() *TileMap {
	tileMap := &TileMap{}

	factory.Init(tileMap)

	return tileMap
}

func (factory TileMapFactory) Init(tileMap *TileMap) {
	SpriteFactory{
		Texture: factory.ComposeTexture(),
	}.Init(&tileMap.Sprite)
	tileMap.SetOriginCenter()
}

func (factory TileMapFactory) ComposeTexture() render.Texture {
	tileMapTexture := factory.CreateTexture()

	tiles := render.TextureFrameExtractor{
		Texture:    factory.TileSet,
		FrameSize:  factory.TileSize,
		FrameCount: factory.TileCount,
	}.Extract()
	transformable := math.Transformable{}

	for rowIndex, row := range factory.Layer {
		for columnIndex, tileValue := range row {
			transformable.SetPosition(math.Vector{
				X: factory.TileSize.X * float64(columnIndex),
				Y: factory.TileSize.Y * float64(rowIndex),
			})

			isTileEmtpy := tileValue == -1

			if isTileEmtpy {
				continue
			}

			tile := tiles[tileValue]

			render.Sprite{
				Texture:   factory.TileSet,
				Rect:      &tile,
				Target:    tileMapTexture,
				Transform: transformable.GetTransform(),
			}.Render()
		}
	}

	return tileMapTexture
}

func (factory TileMapFactory) CreateTexture() render.Texture {
	rows := len(factory.Layer)
	columns := len(factory.Layer[0])

	texture := render.TextureFactory{
		TargetSize: math.Vector{
			X: factory.TileSize.X * float64(columns),
			Y: factory.TileSize.Y * float64(rows),
		},
	}.Create()

	return texture
}
