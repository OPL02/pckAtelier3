package pckAtelier3

import "math/rand"

func ArrayGenerate(nb int) []int {
	arr := make([]int, nb)
	for i := 0; i < nb; i++ {
		arr[i] = rand.Intn(1001)
	}

	return arr
}
