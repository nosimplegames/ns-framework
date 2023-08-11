package hnbUtils

type List[T any] struct {
	Elements []T

	current      *T
	currentIndex int
	hasReachEnd  bool
}

func (list *List[T]) GetCurrent() *T {
	if list.hasReachEnd {
		return nil
	}

	hasToInitCurrent := list.current == nil

	if hasToInitCurrent {
		hasElements := list.Length() > 0

		if hasElements {
			list.current = &list.Elements[0]
		} else {
			list.hasReachEnd = true
		}
	}

	return list.current
}

func (list *List[T]) Next() {
	list.currentIndex++
	list.hasReachEnd = list.currentIndex >= list.Length()

	if list.hasReachEnd {
		list.current = nil
	} else {
		list.current = &list.Elements[list.currentIndex]
	}
}

func (list List[T]) IsLastElement() bool {
	return list.Length()-1 == list.currentIndex
}

func (list List[T]) Length() int {
	return len(list.Elements)
}

func (list List[T]) HasReachEnd() bool {
	return list.hasReachEnd
}
