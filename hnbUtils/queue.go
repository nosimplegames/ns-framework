package hnbUtils

type Queue[T any] struct {
	Elements []T
}

func (queue *Queue[T]) Enqueue(element T) {
	queue.Elements = append(queue.Elements, element)
}

func (queue *Queue[T]) Dequeue() *T {
	hasElement := len(queue.Elements) > 0

	if !hasElement {
		return nil
	}

	element := queue.Elements[0]
	queue.Elements = queue.Elements[1:]
	return &element
}

func (queue Queue[T]) IsEmpty() bool {
	return len(queue.Elements) == 0
}
