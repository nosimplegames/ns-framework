package hnbCore

import (
	"github.com/nosimplegames/ns-framework/hnbNodes"
	"github.com/nosimplegames/ns-framework/hnbPhysics"
	"github.com/nosimplegames/ns-framework/hnbUtils"
)

type EntityChild = hnbNodes.INode[IEntity]
type EntityChildren = []EntityChild
type EntityAdder struct {
	Parent   EntityChild
	Child    EntityChild
	Children EntityChildren
}

func (adder EntityAdder) Add() {
	isParentNil := hnbUtils.IsNil(adder.Parent)

	if isParentNil {
		return
	}

	hnbNodes.ChildAdder[IEntity]{
		Parent:   adder.Parent,
		Child:    adder.Child,
		Children: adder.Children,
	}.Add()

	addChildToParentDangerously(adder.Child, adder.Parent)

	for _, child := range adder.Children {
		addChildToParentDangerously(child, adder.Parent)
	}
}

func addChildToParentDangerously(child EntityChild, parent EntityChild) {
	canAddChild := !hnbUtils.IsNil(child)

	if canAddChild {
		hnbPhysics.AddCollisionable(child.(IEntity))
	}
}
