package core

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
)

func TestRandSlice(testingT *testing.T) {
	randSlice := NewRandSlice(util.RandSeed())

	t := gt.CreateTest(testingT)
	t.Describe("StrElem", func() {
		slice := []string{"abc", "efg", "hij", "klm"}
		t.It("should return a random string element from the slice", func() {
			r := randSlice.StrElem(slice)
			gt.Expect(t, &r).ToBeIn(slice)
		})
	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("IntElem", func() {
		slice := []int{123, 456, 789, 101, 102, 103}
		t2.It("should return a random int element from the slice", func() {
			r := randSlice.IntElem(slice)
			gt.Expect(t2, &r).ToBeIn(slice)
		})
	})
}
