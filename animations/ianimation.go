package animations

type IAnimation interface {
	Update()
	IsAlive() bool
	IsRunning() bool
	Copy() IAnimation
	Reverse() IAnimation
}
