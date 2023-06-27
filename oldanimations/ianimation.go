package oldanimations

type IAnimation interface {
	Update()
	IsAlive() bool
	IsRunning() bool
	Copy() IAnimation
	Reverse() IAnimation
	Pause()
	Resume()
	Die()
}
