package main

import (
	"fmt"
	"time"

	"github.com/yrichika/gfaker/pkg/fk"
)

func main() {
	f := fk.Create()

	// bool
	f.Rand.Bool.Evenly()
	f.Rand.Bool.WeightedToTrue(0.8)

	// int
	f.Rand.Num.IntBt(1, 10)
	f.Rand.Num.Int32Bt(1, 10)
	f.Rand.Num.Int64Bt(1, 10)
	f.Rand.Num.Float32Bt(1.0, 10.0)
	f.Rand.Num.Float64Bt(1.0, 10.0)
	f.Rand.Num.Int()
	f.Rand.Num.Intn(10)

	// string
	f.Rand.Str.Char()
	f.Rand.Str.Letter()
	f.Rand.Str.Digit()
	f.Rand.Str.AlphaRange(5, 10)
	f.Rand.Str.AlphaFixedLength(10)
	f.Rand.Str.AlphaDigitsLike("###-???-***")

	// time
	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)
	f.Rand.Time.PastFuture()
	f.Rand.Time.PastFrom(past30Years)
	f.Rand.Time.Past()
	f.Rand.Time.FutureTo(future30Years)
	f.Rand.Time.Future()
	f.Rand.Time.TimeRange(past30Years, future30Years)
	f.Rand.Time.Duration()
	f.Rand.Time.DurationMilliSec()
	f.Rand.Time.DurationMin()
	f.Rand.Time.DurationHour()
	f.Rand.Time.DurationTo(1 * time.Second)
	f.Rand.Time.DurationRange(1*time.Second, 2*time.Second)

	// slice
	f.Rand.Slice.IntElem([]int{1, 2, 3})
	f.Rand.Slice.StrElem([]string{"foo", "bar", "bazz"})

	// map
	simpleValues := map[any]any{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
		"key4": "value4",
	}
	f.Rand.Map.KeyValue(simpleValues)

	sliceValues := map[any][]any{
		1: {"value11", "value12"},
		2: {"value21", "value22"},
	}
	f.Rand.Map.KeySliceValue(sliceValues)

	// barcode
	f.Barcode.Ean8()
	f.Barcode.Ean13()
	f.Barcode.Isbn10()
	f.Barcode.Isbn13()

	// color
	f.Color.SafeName()
	f.Color.Name()
	f.Color.Hex()
	f.Color.SafeHex()
	f.Color.RgbAsNum()
	f.Color.RgbAsStr()
	f.Color.RgbAsArr()
	f.Color.RgbCss()
	f.Color.RgbaCss()
	f.Color.HslAsNum()
	f.Color.HslAsStr()
	f.Color.HslAsArr()

	//  file
	f.File.MimeType()
	f.File.Extension()

	// destDir := "./tmp"
	// content := "Hello, World!"
	// returnFullPath := false
	// f.File.WriteWithText(destDir, content, "txt", returnFullPath)

	// srcFilePath := "./file/sample.txt"
	// f.File.CopyFrom(destDir, srcFilePath, "txt", returnFullPath)

	// internet
	f.Internet.UserName()
	f.Internet.DomainWord()
	f.Internet.Tld()
	f.Internet.DomainName()
	f.Internet.Email()
	f.Internet.Password()
	f.Internet.Ipv4()
	f.Internet.Ipv6()
	f.Internet.LocalIpv4()
	f.Internet.MacAddress()

	// lorem
	// f.Lorem.Word()
	fmt.Printf("%s\n", r)

	// f.Lorem.WordSliceFixedLength(5)

	// f.Lorem.WordSlice(5)

	// f.Lorem.Words(5)

	// f.Lorem.SentenceFixedLength(5)

	// f.Lorem.Sentence(5)

	// f.Lorem.SentenceSliceFixedLength(5, 5)

	// f.Lorem.SentenceSlice(5, 5)

	// f.Lorem.Sentences(5, 5)

	// f.Lorem.ParagraphSliceFixedLength(5, 5)

	// f.Lorem.ParagraphSlice(5, 5)

	// f.Lorem.Paragraphs(5, 5)

	// // address
	// f.Address.CitySuffix()

	// f.Address.CityPrefix()

	// f.Address.CityName()

	// f.Address.City()

	// f.Address.StreetSuffix()

	// f.Address.StreetName()

	// f.Address.Street()

	// f.Address.BuildingNumber()

	// f.Address.SecondaryAddress()

	// f.Address.StreetAddress()

	// f.Address.Postcode()

	// f.Address.StateAbbr()

	// f.Address.State()

	// f.Address.Address()

	// f.Address.Country()

	// f.Address.Prefecture()

	// f.Address.WardSuffix()

	// f.Address.WardName()

	// f.Address.Ward()

	// f.Address.AreaName()

	// f.Address.AreaNumber()

	// f.Address.Area()

	// f.Address.BuildingName()

	// f.Address.RoomNumber()

	// f.Address.Latitude()

	// f.Address.Longitude()

	// f.Address.LocalCoordinates()

	// // company
	// f.Company.CompanyName()

	// f.Company.CompanyPrefix()

	// f.Company.CompanySuffix()

	// f.Company.Name()

	// f.Company.JobTitleName()

	// f.Company.JobTitle()

	// f.Company.EinPrefix()

	// f.Company.Ein()

	// // person

	// f.Person.FirstNameMale()

	// f.Person.FirstNameFemale()

	// f.Person.FirstName()

	// f.Person.LastName()

	// f.Person.TitleMale()

	// f.Person.TitleFemale()

	// f.Person.Title()

	// f.Person.Suffix()

	// f.Person.MaleName()

	// f.Person.FemaleName()

	// f.Person.Name()

	// f.Person.Ssn()

	// f.Person.FirstKanaNameMale()

	// f.Person.FirstKanaNameFemale()

	// f.Person.FirstKanaName()

	// f.Person.LastKanaName()

	// f.Person.MaleKanaName()

	// f.Person.FemaleKanaName()

	// f.Person.KanaName()

}
