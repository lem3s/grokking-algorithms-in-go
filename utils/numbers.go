package utils

import "math/rand/v2"

func generateInRandomOrder(initial, final int) []int {
	var res []int

	for i := initial; i <= final; i++ {
		res = append(res, i)
	}

	scrambleSlice(res)

	return res
}

func generateInOrder(initial, final int) []int {
	var res []int

	for i := initial; i <= final; i++ {
		res = append(res, i)
	}
	return res
}

func swap(a, b *int) {
	temp := a
	a = b
	b = temp
}

func scrambleSlice(slice []int) []int {
	randomIndexInSlice := rand.IntN(len(slice))

	for number := range slice {
		swap(&number, &slice[randomIndexInSlice])
	}

	return slice
}
