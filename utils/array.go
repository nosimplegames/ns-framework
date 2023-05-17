package utils

type GenericCallback[T interface{}] func(T)

func ForEachBackwards[T interface{}](array []T, callback GenericCallback[T]) {
	for i := len(array) - 1; i >= 0; i-- {
		callback(array[i])
	}
}

type ReduceCallback[T any, K any] func(K, T) K

func Reduce[T any, K any](array []T, callback ReduceCallback[T, K], initialValue K) K {
	accumulator := initialValue

	for _, element := range array {
		accumulator = callback(accumulator, element)
	}

	return accumulator
}
