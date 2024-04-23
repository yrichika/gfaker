package fk

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
)

func TestFaker(testingT *testing.T) {
	t := gt.CreateTest(testingT)

	t.Describe("just tinkering", func() {
		// f := Create()
		j := ja_JP.New()
		f := CreateWithLocale(j)
		t.Test("Faker should ...", func() {
			f.Rand.Bool.Evenly()
			f.Rand.Bool.WeightedToTrue(0.8)

			f.Rand.Num.IntBt(1, 10)
			f.Rand.Num.Int32Bt(1, 10)
			f.Rand.Num.Int64Bt(1, 10)
			f.Rand.Num.Float32Bt(1.0, 10.0)
			f.Rand.Num.Float64Bt(1.0, 10.0)

			// rand.Randのメソッドを使いたい場合は、エイリアスが用意されています
			f.Rand.Num.Int()
			f.Rand.Num.Intn(10)
			f.Rand.Str.Char()
			f.Rand.Time.PastFuture()

			f.Rand.Slice.IntElem([]int{1, 2, 3})
			f.Barcode.Ean13()
			f.Color.Hex()
			f.File.MimeType()
			f.Image.Binary(100, 100, "jpg")
			f.Internet.FirstName()
			f.Lorem.Word()
			f.Address.City()
			f.Company.Name()
		})

	})
}
