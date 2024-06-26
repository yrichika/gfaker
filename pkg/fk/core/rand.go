package core

import (
	"math/rand"
)

type Rand struct {
	Str   *RandStr
	Num   *RandNum
	Bool  *RandBool
	Slice *RandSlice
	Map   *RandMap
	Time  *RandTime
}

func NewRand(rand *rand.Rand) *Rand {
	return &Rand{
		Str:   NewRandStr(rand),
		Num:   NewRandNum(rand),
		Bool:  NewRandBool(rand),
		Slice: NewRandSlice(rand),
		Map:   NewRandMap(rand),
		Time:  NewRandTime(rand),
	}
}
