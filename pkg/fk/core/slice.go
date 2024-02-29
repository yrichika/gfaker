package core

import (
	"math/rand"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

type RandSlice struct {
	rand *rand.Rand
}

func NewRandSlice(rand *rand.Rand) *RandSlice {
	return &RandSlice{
		rand,
	}
}

func (r *RandSlice) StrElem(slice []string) string {
	if len(slice) == 0 {
		log.WrongUsage("Given slice is empty.", 1)
		return ""
	}
	return slice[r.rand.Intn(len(slice))]
}

func (r *RandSlice) IntElem(slice []int) int {
	if len(slice) == 0 {
		log.WrongUsage("Given slice is empty.", 1)
		return 0
	}
	return slice[r.rand.Intn(len(slice))]
}
