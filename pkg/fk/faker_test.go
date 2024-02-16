package fk

import (
	"fmt"
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
)

func TestFaker(testingT *testing.T) {
	t := gt.CreateTest(testingT)

	t.Describe("just tinkering", func() {
		f := Create()
		j := ja_JP.New()
		jp := CreateWithLocale(j)
		t.It("should ...", func() {
			// fmt.Println(f.Str.Char())
			// fmt.Println(f.Rand.Str.AlphaRange(3, 3))
			// fmt.Println(f.Rand.Num.Int(3, -1))
			// fmt.Println(f.Bool.Evenly())

			// fmt.Println(f.Rand.Bool.WeightedToTrue(0.5))
			// fmt.Println(f.Person.FirstNameMale())
			fmt.Println(f.Person.Name())
			fmt.Println(jp.Person.Name())
			fmt.Println(f.Person.FirstKanaNameMale())
			// だいたいこんな感じで
			// f.Person.FirstName()
			// f.Person.LastName()
			// f.Location.City()
			// f.Location.Country()
			// f.Location.Postcode()
			// f.Location.State()
			// f.Location.Street()
			// f.Location.StreetAddress()
			// f.Location.TimeZone()
			// f.Location.Latitude()
			// f.Location.Longitude()
		})
	})
}
