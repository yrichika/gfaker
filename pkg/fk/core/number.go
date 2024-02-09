package core

import (
	"log"
	"math/rand"

	"github.com/yrichika/gfaker/pkg/fk/common"
)

type RandNum struct {
	rand *rand.Rand
}

func NewRandNum(rand *rand.Rand) *RandNum {
	return &RandNum{
		rand,
	}
}

func (r *RandNum) Int(min int, max int) int {
	if min > max {
		file, line := common.GetCallerInfo(1)
		log.Printf("Error at %s: line %d: Invalid range: min=%d, max=%d", file, line, min, max)
		return 0
	}
	return r.rand.Intn(max-min) + min
}

// floatやcomplexなども作る
