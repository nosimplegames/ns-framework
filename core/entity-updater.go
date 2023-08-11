package core

type EntityUpdater struct {
	Entity IEntity
}

func (updater EntityUpdater) UpdateFrame() {
	updater.handleInput()
	updater.updateEntity()
}

func (updater EntityUpdater) handleInput() {
	updater.Entity.HandleInput()

	for _, child := range updater.Entity.GetChildren() {
		EntityUpdater{
			Entity: child.(IEntity),
		}.handleInput()
	}
}

func (updater EntityUpdater) updateEntity() {
	updater.Entity.UpdateFrame()

	for _, child := range updater.Entity.GetChildren() {
		EntityUpdater{
			Entity: child.(IEntity),
		}.updateEntity()
	}

	updater.Entity.RemoveDeadChildren()
}
