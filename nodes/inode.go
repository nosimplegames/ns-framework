package nodes

type INode[T any] interface {
	ILiving

	GetId() string
	SetId(string)
	GetType() string
	SetType(string)

	GetChildren() []T
	RemoveDeadChildren()
	SetParent(T)
}
