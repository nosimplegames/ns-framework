package hnbEntities

import (
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type TileMap struct {
	Sprite
}

type TileMapFactory struct {
	TileSet   hnbRender.Texture
	TileSize  hnbMath.Vector
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

func (factory TileMapFactory) ComposeTexture() hnbRender.Texture {
	tileMapTexture := factory.CreateTexture()

	tiles := hnbRender.TextureFrameExtractor{
		Texture:    factory.TileSet,
		FrameSize:  factory.TileSize,
		FrameCount: factory.TileCount,
	}.Extract()
	transformable := hnbMath.Transformable{}

	for rowIndex, row := range factory.Layer {
		for columnIndex, tileValue := range row {
			transformable.SetPosition(hnbMath.Vector{
				X: factory.TileSize.X * float64(columnIndex),
				Y: factory.TileSize.Y * float64(rowIndex),
			})

			isTileEmtpy := tileValue == -1

			if isTileEmtpy {
				continue
			}

			tile := tiles[tileValue]

			hnbRender.Sprite{
				Texture:   factory.TileSet,
				Rect:      &tile,
				Target:    tileMapTexture,
				Transform: transformable.GetTransform(),
			}.Render()
		}
	}

	return tileMapTexture
}

func (factory TileMapFactory) CreateTexture() hnbRender.Texture {
	rows := len(factory.Layer)
	columns := len(factory.Layer[0])

	texture := hnbRender.TextureFactory{
		TargetSize: hnbMath.Vector{
			X: factory.TileSize.X * float64(columns),
			Y: factory.TileSize.Y * float64(rows),
		},
	}.Create()

	return texture
}
