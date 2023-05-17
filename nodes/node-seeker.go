package nodes

type NodeSeeker[T INode[T]] struct {
	Node INode[T]
}

func (seeker NodeSeeker[T]) IsThereOfType(nodeType string) bool {
	hasSameType := seeker.Node.GetType() == nodeType

	if hasSameType {
		return true
	}

	for _, child := range seeker.Node.GetChildren() {
		childSeeker := NodeSeeker[T]{
			Node: child,
		}

		hasSameType = childSeeker.IsThereOfType(nodeType)
		if hasSameType {
			return true
		}
	}

	return false
}
