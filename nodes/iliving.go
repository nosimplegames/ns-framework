package nodes

type ILiving interface {
	IsAlive() bool
	Die()

	IsRunning() bool
	Pause()
	Resume()
}
