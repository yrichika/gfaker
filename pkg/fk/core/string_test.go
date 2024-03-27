package core

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestRandStr(testingT *testing.T) {
	randStr := NewRandStr(util.RandSeed())

	tChar := gt.CreateTest(testingT)
	tChar.Describe("Char", func() {
		tChar.It("should return a single char includes alpha, num, special", func() {
			r := randStr.Char()
			// check output
			testutil.Output("RandStr.Char", r)
		})
	})

	tLet := gt.CreateTest(testingT)
	tLet.Describe("Letter", func() {
		tLet.It("should return a single letter", func() {
			r := randStr.Letter()
			// check output
			testutil.Output("RandStr.Letter", r)
		})
	})

	tDigit := gt.CreateTest(testingT)
	tDigit.Describe("Digit", func() {
		tDigit.It("should return a single digit", func() {
			r := randStr.Digit()
			// check output
			testutil.Output("RandStr.Digit", r)
		})
	})

	tAlphaRange := gt.CreateTest(testingT)
	tAlphaRange.Describe("AlphaRange", func() {
		tAlphaRange.It("should return a string with random length", func() {
			r := randStr.AlphaRange(1, 5)
			length := len(r)
			gt.Expect(tAlphaRange, &length).ToBe_(gt.Between(1, 5))
			// check output
			testutil.Output("RandStr.AlphaRange", r)
		})
	})

	tAlphaFixed := gt.CreateTest(testingT)
	tAlphaFixed.Describe("AlphaFixedLength", func() {
		tAlphaFixed.It("should return a string with fixed length", func() {
			r := randStr.AlphaFixedLength(5)
			length := len(r)
			gt.Expect(tAlphaFixed, &length).ToBe(5)
			// check output
			testutil.Output("RandStr.AlphaFixedLength", r)
		})
	})

	tAlDgLike := gt.CreateTest(testingT)
	tAlDgLike.Describe("AlphaDigitsLike", func() {
		tAlDgLike.It("should return a string with specified alpha and digits", func() {
			r := randStr.AlphaDigitsLike("abc-###-???")
			gt.Expect(tAlDgLike, &r).ToMatchRegex("abc-[0-9]{3}-[a-zA-Z]{3}")
		})
	})
}
