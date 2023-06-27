package nodes

import "github.com/nosimplegames/ns-framework/utils"

type Node[T INode[T]] struct {
	Living

	id    string
	type_ string

	children []INode[T]
	parent   INode[T]
}

func (node *Node[T]) addChild(child INode[T]) {
	node.children = append(node.children, child)
}

func (node Node[T]) GetChildren() []INode[T] {
	return node.children
}

func (node *Node[T]) setParent(parent INode[T]) {
	node.parent = parent
}

func (node Node[T]) GetParent() (INode[T], bool) {
	return node.parent, node.parent != nil
}

func (node *Node[T]) RemoveDeadChildren() {
	utils.RemoveDead(&node.children)
}

func (node *Node[T]) RemoveChildren() {
	for _, child := range node.children {
		child.RemoveChildren()
		child.Die()
	}

	node.children = []INode[T]{}
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
