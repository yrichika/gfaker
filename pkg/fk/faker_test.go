package fk

import (
	"fmt"
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
)

func TestFake(testingT *testing.T) {
	t := gt.CreateTest(testingT)

	t.Describe("just tinkering", func() {
		f := Create()
		j := ja_JP.New()
		jp := CreateWithLocale(j)
		t.It("should ...", func() {
			// fmt.Println(f.Str.Char())
			// fmt.Println(f.Str.AlphaRange(3, 10))
			// fmt.Println(f.Num.Int(3, 10))
			// fmt.Println(f.Bool.Evenly())

			// fmt.Println(f.Rand.Bool.WeightedToTrue(0.5))
			// fmt.Println(f.Person.FirstNameMale())
			fmt.Println(f.Person.Name())
			fmt.Println(jp.Person.Name())
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
