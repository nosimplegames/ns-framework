package hnbCore

import "github.com/nosimplegames/ns-framework/hnbUtils"

type Updatables struct {
	Updatables []IUpdatable
}

func (updatables *Updatables) AddUpdatable(updatable IUpdatable) {
	updatables.Updatables = append(updatables.Updatables, updatable)
}

func (updatables *Updatables) UpdateFrame() {
	for _, updatable := range updatables.Updatables {
		updatable.UpdateFrame()
	}

	hnbUtils.RemoveDead(&updatables.Updatables)
}

var updatables *Updatables = nil

func GetUpdatables() *Updatables {
	mustInitUpdatables := updatables == nil

	if mustInitUpdatables {
		updatables = &Updatables{}
	}

	return updatables
}
