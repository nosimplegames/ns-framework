package hnbNodes

import "github.com/nosimplegames/ns-framework/hnbUtils"

type ChildAdder[T INode[T]] struct {
	Parent   INode[T]
	Child    INode[T]
	Children []INode[T]
}

func (adder ChildAdder[T]) Add() {
	adder.addChildToParent(
		adder.Child,
		adder.Parent,
	)

	for _, child := range adder.Children {
		adder.addChildToParent(child, adder.Parent)
	}
}

func (adder ChildAdder[T]) addChildToParent(child INode[T], parent INode[T]) {
	canAdd := !hnbUtils.IsNil(child) && !hnbUtils.IsNil(parent)

	if !canAdd {
		return
	}

	parent.addChild(child)
	child.setParent(parent)
}
