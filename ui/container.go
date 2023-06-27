package ui

// import (
// 	"github.com/nosimplegames/ns-framework/core"
// 	"github.com/nosimplegames/ns-framework/math"
// 	"github.com/nosimplegames/ns-framework/ui/uicontainer"
// )

// type Container struct {
// 	Component

// 	Padding Padding
// 	Layout  ILayout

// 	sizeCalculationPolicy  uicontainer.SizeCalculationPolicy
// 	needToPositionElements bool
// 	needToCalculateSize    bool
// }

// func (container *Container) Update() {
// 	hasLayout := container.Layout != nil

// 	if !hasLayout {
// 		return
// 	}

// 	container.RepositionElements()
// }

// func (container *Container) RepositionElements() {
// 	if !container.needToPositionElements {
// 		return
// 	}

// 	container.Layout.SetSize(container.GetLayoutSize())
// 	positions := container.Layout.GetPositions(container.GetElements())

// 	for childIndex, child := range container.GetChildren() {
// 		position := positions[childIndex]
// 		child.SetPosition(position)
// 	}

// 	container.needToPositionElements = false
// }

// func (container *Container) AddChild(child core.IEntity) {
// 	container.Component.AddChild(child)

// 	container.needToPositionElements = true
// 	container.needToCalculateSize = true

// 	component, isComponent := child.(*Component)

// 	if !isComponent {
// 		return
// 	}

// 	container.connectToComponent(component)
// }

// func (container *Container) connectToComponent(component *Component) {
// 	component.OnSizeChanged.AddCallback(func() {
// 		container.needToPositionElements = true
// 		container.needToCalculateSize = true
// 	})
// }

// func (container *Container) GetElements() []LayoutElement {
// 	elements := []LayoutElement{}

// 	for _, child := range container.GetChildren() {
// 		elements = append(elements, LayoutElement{
// 			Size: child.GetSize(),
// 		})
// 	}

// 	return elements
// }

// func (container *Container) SetSizeCalculationPolicy(policy uicontainer.SizeCalculationPolicy) {
// 	container.sizeCalculationPolicy = policy
// }

// func (container *Container) SetSize(size math.Vector) {
// 	container.needToCalculateSize = false
// 	container.SetSizeCalculationPolicy(uicontainer.Manual)
// 	container.Component.SetSize(size)
// }

// func (container *Container) GetSize() math.Vector {
// 	needToCalculateSize := container.sizeCalculationPolicy == uicontainer.Auto && container.needToCalculateSize

// 	if !needToCalculateSize {
// 		return container.Component.GetSize()
// 	}

// 	size := container.calculateContentSize()
// 	container.Component.SetSize(size)
// 	return size
// }

// func (container Container) GetLayoutSize() math.Vector {
// 	size := container.GetSize()

// 	layoutSize := math.Vector{
// 		X: size.X - container.Padding.Right - container.Padding.Left,
// 		Y: size.Y - container.Padding.Top - container.Padding.Bottom,
// 	}

// 	return layoutSize
// }

// func (container Container) calculateContentSize() math.Vector {
// 	size := container.Layout.GetContentSize(container.GetElements())
// 	size.X += container.Padding.Right + container.Padding.Left
// 	size.Y += container.Padding.Top + container.Padding.Bottom

// 	return size
// }
