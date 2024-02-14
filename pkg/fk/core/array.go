package core

import "math/rand"

type RandArray struct {
	rand *rand.Rand
}

func NewRandArray(rand *rand.Rand) *RandArray {
	return &RandArray{
		rand,
	}
}

func (r *RandArray) StrElem(arr []string) string {
	return arr[r.rand.Intn(len(arr))]
}

func (r *RandArray) IntElem(arr []int) int {
	return arr[r.rand.Intn(len(arr))]
}
