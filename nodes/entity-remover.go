package nodes

type NodeRemover[T INode[T]] struct {
	Parent INode[T]
	Id     string
}

func (remover NodeRemover[T]) Remove() bool {
	for _, child := range remover.Parent.GetChildren() {
		isIt := child.GetId() == remover.Id

		if isIt {
			child.Die()
			return true
		}
	}

	wasRemovedFromChildren := false

	for _, child := range remover.Parent.GetChildren() {
		wasRemovedFromChildren = NodeRemover[T]{
			Parent: child,
			Id:     remover.Id,
		}.Remove()

		if wasRemovedFromChildren {
			break
		}
	}

	return wasRemovedFromChildren
}
