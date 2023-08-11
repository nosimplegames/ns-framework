package hnbCore

import (
	"errors"
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/nosimplegames/ns-framework/hnbMath"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"github.com/nosimplegames/ns-framework/hnbRender"
)

type Game struct {
	ScenesManager

	BackgroundColor color.Color
	Cameras         []ICamera
	WindowSize      hnbMath.Vector
	Size            hnbMath.Vector

	MustPrintDebug bool
	MustDrawWorld  bool
}

func (game *Game) Update(*ebiten.Image) error {
	GetUpdatables().UpdateFrame()
	GetAnimations().UpdateFrame()
	game.UpdateEntities()
	hnbPhysics.GetWorld().UpdateFrame()

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
	scene, hasScene := game.GetScene()

	if !hasScene {
		return
	}

	EntityUpdater{
		Entity: scene,
	}.UpdateFrame()
}

func (game Game) Draw(screen hnbRender.RenderTarget) {
	hasCameras := len(game.Cameras) > 0

	if hasCameras {
		game.DrawCameras(screen)
	} else {
		game.DrawEntities(screen)
	}
}

func (game Game) DrawCameras(screen hnbRender.RenderTarget) {
	for _, camera := range game.Cameras {
		cameraTexture, _ := ebiten.NewImage(int(game.Size.X), int(game.Size.Y), ebiten.FilterDefault)
		cameraTexture.Fill(game.BackgroundColor)
		game.DrawEntities(cameraTexture)

		camera.SetTexture(cameraTexture)
		camera.Draw(screen, camera.GetTransform())
	}
}

func (game Game) DrawEntities(target hnbRender.RenderTarget) {
	scene, hasScene := game.GetScene()

	if !hasScene {
		return
	}

	EntityDrawer{
		Entity:    scene,
		Transform: hnbMath.Transform{},
		Target:    target,
	}.Draw()

	if game.MustPrintDebug {
		totalEntities := len(scene.GetChildren())
		ebitenutil.DebugPrint(target, strconv.Itoa(totalEntities))
	}

	if game.MustDrawWorld {
		world := hnbPhysics.GetWorld()
		world.Draw(target)
	}
}

func (game *Game) PushScene(scene IScene) {
	game.ScenesManager.PushScene(scene)

	currentScene, hasScene := game.GetScene()

	if !hasScene {
		return
	}

	root = currentScene
}

func (game Game) Layout(_, _ int) (int, int) {
	return int(game.WindowSize.X), int(game.WindowSize.Y)
}

func (game *Game) Run() {
	ebiten.SetWindowSize(int(game.WindowSize.X), int(game.WindowSize.Y))
	game.SetDefaultBackgroundColor()
	ebiten.RunGame(game)
}

func (game *Game) SetDefaultBackgroundColor() {
	hasBackgroundColor := game.BackgroundColor != nil

	if !hasBackgroundColor {
		game.BackgroundColor = color.Black
	}
}

func (game Game) GetDefaultRenderingBox() hnbMath.Box {
	return hnbMath.Box{
		Position: game.Size.By(0.5),
		Size:     game.Size,
	}
}

func (game Game) GetDefaultViewport() hnbMath.Box {
	return hnbMath.Box{
		Position: game.WindowSize.By(0.5),
		Size:     game.WindowSize,
	}
}

func (game *Game) AddCamera(camera ICamera) {
	game.Cameras = append(game.Cameras, camera)
}

var root IEntity

func GetRoot() IEntity {
	return root
}

func AddChildToRoot(child IEntity) {
	EntityAdder{
		Parent: root,
		Child:  child,
	}.Add()
}
