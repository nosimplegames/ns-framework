package core

type ScenesManager struct {
	scenes       []IScene
	currentScene IScene
}

func (manager *ScenesManager) PushScene(scene IScene) {
	manager.scenes = append(manager.scenes, scene)
	manager.currentScene = scene
}

func (manager *ScenesManager) GetScene() (IScene, bool) {
	return manager.currentScene, manager.currentScene != nil
}
