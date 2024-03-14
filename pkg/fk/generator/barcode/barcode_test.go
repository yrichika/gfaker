package barcode

import (
	"strconv"
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
)

func TestBarcode(testingT *testing.T) {

	coreRand := core.NewRand(util.RandSeed())
	barcode := New(coreRand)

	tEan := gt.CreateTest(testingT)
	tEan.Describe("EAN", func() {

		tEan.Test("Ean8 should return a barcode with 8 digits", func() {
			r := barcode.Ean8()
			gt.Expect(tEan, &r).ToMatchRegex(`^\d{8}$`)

			lastDigit, _ := strconv.Atoi(r[7:])
			heading7Digits := r[:7]
			gt.Expect(tEan, &lastDigit).ToBe(calcEanCheckDigit(heading7Digits))
		})

		tEan.Test("Ean13 should return a barcode with 13 digits", func() {
			r := barcode.Ean13()
			gt.Expect(tEan, &r).ToMatchRegex(`^\d{13}$`)

			lastDigit, _ := strconv.Atoi(r[12:])
			heading12Digits := r[:12]
			gt.Expect(tEan, &lastDigit).ToBe(calcEanCheckDigit(heading12Digits))
		})

		tEan.Test("checkDigit should calculate barcode check digit", func() {
			testCases := []struct {
				digits string
				want   int
			}{
				// 7 digits -> EAN-8
				{"1234567", 0},
				{"7654321", 0},
				{"3897546", 2},
				{"7653573", 4},
				{"3264902", 4},
				// 12 digits -> EAN-13
				{"764564239870", 7},
				{"233209246782", 8},
				{"876456876876", 6},
				{"272549071238", 4},
				{"986126758742", 7},
			}

			for _, tc := range testCases {
				result := calcEanCheckDigit(tc.digits)
				gt.Expect(tEan, &result).ToBe(tc.want)
			}
		})
	})

	tIsbn := gt.CreateTest(testingT)
	tIsbn.Describe("ISBN", func() {
		tIsbn.Test("Isbn10 should return a ISBN with 10 digits", func() {
			r := barcode.Isbn10()
			gt.Expect(tIsbn, &r).ToMatchRegex(`^\d{9}[\dX]$`)

			lastDigit := r[9:]
			heading9Digits := r[:9]
			gt.Expect(tEan, &lastDigit).ToBe(calcIsbnCheckDigit(heading9Digits))
		})

		tIsbn.Test("Isbn13 should return a ISBN with 13 digits", func() {
			r := barcode.Isbn13()
			gt.Expect(tIsbn, &r).ToMatchRegex(`^97[89]\d{10}$`)

			lastDigit, _ := strconv.Atoi(r[12:])
			heading12Digits := r[:12]
			gt.Expect(tEan, &lastDigit).ToBe(calcEanCheckDigit(heading12Digits))
		})

		tIsbn.Test("calcIsbnCheckDigit should calculate ISBN check digit", func() {
			testCases := []struct {
				input string
				want  string
			}{
				{"765757657", "X"},
				{"111111111", "1"},
				{"764350980", "8"},
				{"325468122", "3"},
				{"753697616", "X"},
				{"456893287", "4"},
			}

			for _, tc := range testCases {
				got := calcIsbnCheckDigit(tc.input)
				gt.Expect(tIsbn, &got).ToBe(tc.want)
			}
		})
	})
}
