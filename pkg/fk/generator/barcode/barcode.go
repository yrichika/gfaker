package barcode

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/yrichika/gfaker/pkg/fk/core"
)

type Barcode struct {
	rand *core.Rand
}

func New(rand *core.Rand) *Barcode {
	return &Barcode{
		rand,
	}
}

// example: '12345678'
func (b *Barcode) Ean8() string {
	return b.ean(8)
}

// example: '1234567890123'
func (b *Barcode) Ean13() string {
	return b.ean(13)
}

func (b *Barcode) ean(length int) string {
	if length != 8 && length != 13 {
		panic(fmt.Sprintf("Invalid length for EAN barcode. EAN code length must be 8 or 13. Given length: [%d]", length))
	}
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, length-1))
	code := b.rand.Str.AlphaDigitsLike(format)
	return code + fmt.Sprint(calcEanCheckDigit(code))
}

// Checksum computes the checksum of an EAN number.
// See: https://en.wikipedia.org/wiki/International_Article_Number
func calcEanCheckDigit(digits string) int {
	sequence := []int{1, 3}
	if len(digits)+1 == 8 {
		sequence = []int{3, 1}
	}
	sums := 0

	for n, digit := range digits {
		sums += int(digit-'0') * sequence[n%2]
	}

	return (10 - sums%10) % 10
}

// Get a random ISBN-10 code
// See http://en.wikipedia.org/wiki/International_Standard_Book_Number
// example: '3254681223'
func (b *Barcode) Isbn10() string {
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, 9))
	code := b.rand.Str.AlphaDigitsLike(format)
	return code + fmt.Sprint(calcIsbnCheckDigit(code))
}

// Get a random ISBN-13 code
// See http://en.wikipedia.org/wiki/International_Standard_Book_Number
// example: '9783254681223'
func (b *Barcode) Isbn13() string {
	digitSym := []byte("#")
	format := string(bytes.Repeat(digitSym, 9))
	prefX := b.rand.Num.IntBt(8, 9)
	prefix := "97" + fmt.Sprint(prefX)

	pubBookCode := b.rand.Str.AlphaDigitsLike(format)
	code := prefix + pubBookCode
	// EAN check digit is used here because it's 12 digits
	return code + fmt.Sprint(calcEanCheckDigit(code))
}

// Checksum calculates the ISBN-10 check digit
// See: http://en.wikipedia.org/wiki/International_Standard_Book_Number#ISBN-10_check_digits
func calcIsbnCheckDigit(input string) string {
	// We're calculating check digit for ISBN-10
	// so, the length of the input should be 9
	length := 9
	if len(input) != length {
		panic(fmt.Sprintf("input length should be equal to %d", length))
	}

	sum := 0
	for i, digit := range input {
		sum += (10 - i) * int(digit-'0')
	}

	result := (11 - sum%11) % 11

	// 10 is replaced by X
	if result < 10 {
		return strconv.Itoa(result)
	}
	return "X"
}
