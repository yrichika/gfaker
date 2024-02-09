package core

import (
	"math/rand"
)

type Rand struct {
	Str  *RandStr
	Num  *RandNum
	Bool *RandBool
}

func NewRand(rand *rand.Rand) *Rand {
	return &Rand{
		Str:  NewRandStr(rand),
		Num:  NewRandNum(rand),
		Bool: NewRandBool(rand),
	}
}
