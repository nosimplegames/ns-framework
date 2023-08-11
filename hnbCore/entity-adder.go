package hnbCore

import "github.com/nosimplegames/ns-framework/hnbNodes"

type EntityChild = hnbNodes.INode[IEntity]
type EntityChildren = []EntityChild
type EntityAdder = hnbNodes.ChildAdder[IEntity]
