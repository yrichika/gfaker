package util

import (
	"fmt"
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestUtil(testingT *testing.T) {

	rand := RandSeed()

	t := gt.CreateTest(testingT)
	t.Describe("util", func() {
		t.Todo("create tests")
	})

	t2 := gt.CreateTest(testingT)
	t2.Describe("TruncateToPrecision", func() {
		t2.It("should truncate to precision", func() {
			precision := rand.Intn(8)
			r := TruncateToPrecision(1.23456789123456, precision)
			fmt.Println(r)
			decimalLength := testutil.GetDecimalLength(r)
			fmt.Println(decimalLength)
			gt.Expect(t2, &decimalLength).ToBe(precision)
		})
	})

	t3 := gt.CreateTest(testingT)
	t3.Describe("CapFirstLetter", func() {
		t3.It("should capitalize the first letter", func() {
			r := CapFirstLetter("hello")
			gt.Expect(t3, &r).ToBe("Hello")

			r2 := CapFirstLetter("world")
			gt.Expect(t3, &r2).ToBe("World")
		})
	})
}
