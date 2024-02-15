package core

import (
	"fmt"
	"math/rand"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

type RandBool struct {
	rand *rand.Rand
}

func NewRandBool(rand *rand.Rand) *RandBool {
	return &RandBool{
		rand,
	}
}

// 50%の確率でtrueかfalseを返す
func (r *RandBool) Evenly() bool {
	return r.rand.Intn(2) == 0
}

func (r *RandBool) WeightedToTrue(weight float32) bool {
	if weight < 0 || weight > 1 {
		errMsg := fmt.Sprintf("Invalid weight: %f", weight)
		log.WrongUsage(errMsg, 1)
		return false
	}
	return r.rand.Float32() < weight
}
