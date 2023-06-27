package core

import "github.com/nosimplegames/ns-framework/nodes"

type EntityChild = nodes.INode[IEntity]
type EntityChildren = []EntityChild
type EntityAdder = nodes.ChildAdder[IEntity]
