package core

import (
	"errors"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/nosimplegames/ns-framework/math"
	"github.com/nosimplegames/ns-framework/nodes"
	"github.com/nosimplegames/ns-framework/render"
)

type Game struct {
	nodes.Node[IEntity]

	BackgroundColor color.Color
	Cameras         []ICamera
	WindowSize      math.Vector
	Size            math.Vector

	MustPrintDebug bool
	MustDrawWorld  bool
}

func (game *Game) Update(*ebiten.Image) error {
	GetUpdatables().Update()
	GetAnimations().Update()
	game.UpdateEntities()
	GetWorld().Update()

	var err error = nil

	if inpututil.IsKeyJustReleased(ebiten.KeyEscape) {
		err = errors.New("")
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyF1) {
		game.MustPrintDebug = !game.MustPrintDebug
	}

	if inpututil.IsKeyJustReleased(ebiten.KeyF2) {
		game.MustDrawWorld = !game.MustDrawWorld
	}

	return err
}

func (game *Game) UpdateEntities() {
	for _, entity := range game.GetChildren() {
		EntityUpdater{
			Entity: entity,
		}.Update()
	}

	game.RemoveDeadChildren()
}

func (game Game) Draw(screen render.RenderTarget) {
	hasCameras := len(game.Cameras) > 0

	if hasCameras {
		game.DrawCameras(screen)
	} else {
		game.DrawEntities(screen)
	}
}

func (game Game) DrawCameras(screen render.RenderTarget) {
	for _, camera := range game.Cameras {
		cameraTexture, _ := ebiten.NewImage(int(game.Size.X), int(game.Size.Y), ebiten.FilterDefault)
		cameraTexture.Fill(game.BackgroundColor)
		game.DrawEntities(cameraTexture)

		camera.SetTexture(cameraTexture)
		camera.Draw(screen, camera.GetTransform())
	}
}

func (game Game) DrawEntities(target render.RenderTarget) {
	entities := game.GetChildren()
	for _, entity := range entities {
		EntityDrawer{
			Entity:    entity,
			Transform: math.Transform{},
			Target:    target,
		}.Draw()
	}

	if game.MustPrintDebug {
		ebitenutil.DebugPrint(target, strconv.Itoa(len(entities)))
	}

	if game.MustDrawWorld {
		isThereWorld := world != nil

		if isThereWorld {
			world.Draw(target)
		}
	}
}

func (game Game) Layout(_, _ int) (int, int) {
	return int(game.WindowSize.X), int(game.WindowSize.Y)
}

func (game *Game) Run() {
	ebiten.SetWindowSize(int(game.WindowSize.X), int(game.WindowSize.Y))
	game.SetDefaultBackgroundColor()
	// game.InitDefaultCamera()
	ebiten.RunGame(game)
}

func (game *Game) SetDefaultBackgroundColor() {
	hasBackgroundColor := game.BackgroundColor != nil

	if !hasBackgroundColor {
		game.BackgroundColor = color.Black
	}
}

func (game Game) GetDefaultRenderingBox() math.Box {
	return math.Box{
		Position: game.Size.By(0.5),
		Size:     game.Size,
	}
}

func (game Game) GetDefaultViewport() math.Box {
	return math.Box{
		Position: game.WindowSize.By(0.5),
		Size:     game.WindowSize,
	}
}

func (game *Game) AddCamera(camera ICamera) {
	game.Cameras = append(game.Cameras, camera)
}

// func (game *Game) InitDefaultCamera() {
// 	hasCameras := len(game.Cameras) > 0

// 	if hasCameras {
// 		return
// 	}

// 	defaultCamera := &Camera{}
// 	defaultCamera.RenderingBox = math.Box{
// 		Position: game.Size.By(0.5),
// 		Size:     game.Size,
// 	}
// 	defaultCamera.SetViewport(math.Box{
// 		Position: game.WindowSize.By(0.5),
// 		Size:     game.WindowSize,
// 	})
// 	game.Cameras = append(game.Cameras, defaultCamera)
// }

// func (game *Game) GetDefaultCamera() *Camera {
// 	game.InitDefaultCamera()

// 	return game.Cameras[0]
// }
