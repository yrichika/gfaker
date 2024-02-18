package core

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yrichika/gest/pkg/gt"
)

func TestRandNum(testingT *testing.T) {
	t := gt.CreateTest(testingT)

	randNum := NewRandNum(rand.New(rand.NewSource(time.Now().UnixNano())))
	t.Describe("RandNum", func() {
		t.It("should return int value in given range", func() {
			r := randNum.IntBt(1, 10)
			gt.Expect(t, &r).ToBe_(gt.Between(1), 10)
		})
	})
}
