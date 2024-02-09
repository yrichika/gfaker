package core

import (
	"math/rand"

	"log"

	"github.com/yrichika/gfaker/pkg/fk/common"
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
	file, line := common.GetCallerInfo(1)

	if weight < 0 || weight > 1 {
		log.Printf("Error at %s: line %d: Invalid weight: %f", file, line, weight)
		return false
	}
	return r.rand.Float32() < weight
}
