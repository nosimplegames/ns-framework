package core

import "github.com/nosimplegames/ns-framework/utils"

type Updatables struct {
	Updatables []IUpdatable
}

func (updatables *Updatables) AddUpdatable(updatable IUpdatable) {
	updatables.Updatables = append(updatables.Updatables, updatable)
}

func (updatables *Updatables) Update() {
	for _, updatable := range updatables.Updatables {
		updatable.Update()
	}

	utils.RemoveDead(&updatables.Updatables)
}

var updatables *Updatables = nil

func GetUpdatables() *Updatables {
	mustInitUpdatables := updatables == nil

	if mustInitUpdatables {
		updatables = &Updatables{}
	}

	return updatables
}
