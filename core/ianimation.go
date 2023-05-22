package core

type IAnimation interface {
	Update()
	Apply()

	WillReplay() bool
	Restart()

	IsAlive() bool
	IsRunning() bool
	HasFinished() bool
	Pause()
	Resume()
	Stop()

	Copy(interface{}) IAnimation
}
