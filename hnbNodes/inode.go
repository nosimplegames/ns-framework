package hnbNodes

type INode[T any] interface {
	ILiving

	GetId() string
	SetId(string)
	GetType() string
	SetType(string)

	addChild(INode[T])
	GetChildren() []INode[T]
	RemoveDeadChildren()
	RemoveChildren()
	setParent(INode[T])
	GetParent() (INode[T], bool)
}
