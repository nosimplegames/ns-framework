package hnbAssets

import (
	"fmt"

	"golang.org/x/image/font/sfnt"
)

type FontManager struct {
	Fonts map[string]*sfnt.Font
}

type FontData struct {
	Bytes FontBytes
	Name  string
}

func (fontManager *FontManager) GetFont(fontData FontData) *sfnt.Font {
	font := fontManager.Fonts[fontData.Name]
	fontAlreadyLoaded := font != nil

	if !fontAlreadyLoaded {
		font = fontManager.LoadFont(fontData.Bytes)
		fontManager.Fonts[fontData.Name] = font
	}

	return font
}

func (fontManager FontManager) LoadFont(bytes FontBytes) *sfnt.Font {
	font, err := sfnt.Parse(bytes)

	if err != nil {
		fmt.Println("Error getting font")
		panic(err)
	}

	return font
}

var fontManager *FontManager = nil

func GetFontManager() *FontManager {
	needToInitFontManager := fontManager == nil

	if needToInitFontManager {
		fontManager = &FontManager{
			Fonts: map[string]*sfnt.Font{},
		}
	}

	return fontManager
}
