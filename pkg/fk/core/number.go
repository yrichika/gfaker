package core

import (
	"fmt"
	"math/rand"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

type RandNum struct {
	rand *rand.Rand
}

type orderable interface {
	int |
		int8 |
		int16 |
		int32 |
		int64 |
		uint |
		uint8 |
		uint16 |
		uint32 |
		uint64 |
		uintptr |
		float32 |
		float64
}

func NewRandNum(rand *rand.Rand) *RandNum {
	return &RandNum{
		rand,
	}
}

func (r *RandNum) IntBt(min int, max int) int {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Intn(randMax) + min
}

func (r *RandNum) Int32Bt(min int32, max int32) int32 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Int31n(randMax) + min
}

func (r *RandNum) Int64Bt(min int64, max int64) int64 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Int63n(randMax) + min
}

func (r *RandNum) Float32Bt(min float32, max float32) float32 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Float32()*randMax + min
}

func (r *RandNum) Float64Bt(min float64, max float64) float64 {
	randMax, err := randMaxRange(min, max)
	if err != nil {
		return 0
	}
	return r.rand.Float64()*randMax + min
}

// Intn()などに渡すための、ランダム値のmaxの値を返す
// 全ての数値では共通の処理のためまとめているだけ
func randMaxRange[N orderable](min N, max N) (N, error) {
	var err error
	if min >= max {
		errMsg := fmt.Sprintf("Invalid range: min=%v, max=%v", min, max)
		log.WrongUsage(errMsg, 2)
		err = fmt.Errorf(errMsg)
	}

	return max - min, err
}

// inや、floatやcomplexなども作る
