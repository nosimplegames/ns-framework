package animations

import (
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/render"
)

type ISprite interface {
	SetRect(*math.Box)
	SetTexture(render.Texture)
}
