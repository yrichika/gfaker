package core

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestRandArray(testingT *testing.T) {
	randArray := NewRandArray(util.RandSeed())

	t := gt.CreateTest(testingT)
	t.Describe("StrElem", func() {
		arr := []string{"abc", "efg", "hij", "klm"}
		t.It("should return a random string element from the array", func() {
			r := randArray.StrElem(arr)
			// Check with output
			testutil.Output("RandArray.StrElem", r)
		})
	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("IntElem", func() {
		arr := []int{123, 456, 789, 101, 102, 103}
		t2.It("should return a random int element from the array", func() {
			r := randArray.IntElem(arr)
			// Check with output
			testutil.Output("RandArray.IntElem", r)
		})
	})
}
