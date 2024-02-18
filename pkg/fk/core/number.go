package core

import (
	"fmt"
	"math/rand"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

type RandNum struct {
	rand *rand.Rand
}

func NewRandNum(rand *rand.Rand) *RandNum {
	return &RandNum{
		rand,
	}
}

func (r *RandNum) IntBt(min int, max int) int {
	if min > max {
		errMsg := fmt.Sprintf("Invalid range: min=%d, max=%d", min, max)
		log.WrongUsage(errMsg, 1)
		return 0
	}
	return r.rand.Intn(max-min) + min
}

// inや、floatやcomplexなども作る
