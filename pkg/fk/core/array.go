package core

import (
	"math/rand"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

type RandArray struct {
	rand *rand.Rand
}

func NewRandArray(rand *rand.Rand) *RandArray {
	return &RandArray{
		rand,
	}
}

func (r *RandArray) StrElem(arr []string) string {
	if len(arr) == 0 {
		log.WrongUsage("Given array is empty.", 1)
		return ""
	}
	return arr[r.rand.Intn(len(arr))]
}

func (r *RandArray) IntElem(arr []int) int {
	if len(arr) == 0 {
		log.WrongUsage("Given array is empty.", 1)
		return 0
	}
	return arr[r.rand.Intn(len(arr))]
}
