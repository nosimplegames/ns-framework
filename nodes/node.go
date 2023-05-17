package nodes

import "github.com/nosimplegames/ns-framework/utils"

type Node[T INode[T]] struct {
	id    string
	type_ string

	children []T
	parent   T
}

func (node *Node[T]) AddChild(child T) {
	node.children = append(node.children, child)
	child.SetParent(child)
}

func (node *Node[T]) AddChildren(children []T) {
	for _, child := range children {
		node.AddChild(child)
	}
}

func (node Node[T]) GetChildren() []T {
	return node.children
}

func (node Node[T]) SetParent(parent T) {
	node.parent = parent
}

func (node Node[T]) GetParent() T {
	return node.parent
}

func (node *Node[T]) RemoveDeadChildren() {
	utils.RemoveDead(&node.children)
}

func (node Node[T]) SetId(id string) {
	node.id = id
}

func (node Node[T]) GetId() string {
	return node.id
}

func (node Node[T]) SetType(type_ string) {
	node.type_ = type_
}

func (node Node[T]) GetType() string {
	return node.type_
}
