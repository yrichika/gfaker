package core

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
)

func TestRandNum(testingT *testing.T) {
	t := gt.CreateTest(testingT)

	randNum := NewRandNum(util.RandSeed())
	t.Describe("IntBt", func() {
		t.It("should return int value in given range", func() {
			r := randNum.IntBt(1, 10)
			gt.Expect(t, &r).ToBe_(gt.Between(1), 10)
		})

	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("Float32Bt", func() {
		t2.It("should return float value in given range", func() {
			r := randNum.Float32Bt(1.0, 2.0)
			gt.Expect(t2, &r).ToBe_(gt.Between(float32(1.0)), 2.0)
		})
	})
}
