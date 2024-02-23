package core

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestRandBool(testingT *testing.T) {
	randBool := NewRandBool(util.RandSeed())

	t := gt.CreateTest(testingT)
	t.Describe("Evenly", func() {
		t.It("should return true or false", func() {
			r := randBool.Evenly()
			// Check with output
			testutil.Output("RandBool.Evenly", r)
		})
	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("Weighted", func() {
		t2.It("should return weighted bool", func() {
			r := randBool.WeightedToTrue(0.7)
			// Check with output
			testutil.Output("RandBool.WeightedToTrue", r)
		})

		t2.It("should return true when 1", func() {
			r := randBool.WeightedToTrue(1)
			gt.Expect(t2, &r).ToBe(true)
		})

		t2.It("should return false when 0", func() {
			r := randBool.WeightedToTrue(0)
			gt.Expect(t2, &r).ToBe(false)
		})
	})
}
