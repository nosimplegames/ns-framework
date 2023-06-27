package core

type EntityUpdater struct {
	Entity IEntity
}

func (updater EntityUpdater) Update() {
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
	updater.Entity.Update()

	for _, child := range updater.Entity.GetChildren() {
		EntityUpdater{
			Entity: child.(IEntity),
		}.updateEntity()
	}

	updater.Entity.RemoveDeadChildren()
}
