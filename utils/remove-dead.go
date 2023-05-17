package utils

type IBeing interface {
	IsAlive() bool
}

func RemoveDead[T IBeing](beings *[]T) {
	for i := 0; i < len(*beings); {
		entity := (*beings)[i]

		if entity.IsAlive() {
			i++
		} else {
			*beings = append((*beings)[:i], (*beings)[i+1:]...)
		}
	}
}
