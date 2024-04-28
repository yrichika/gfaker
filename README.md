# gfaker

gfakerは、ランダムなダミーデータを作成するためのライブラリです。
テストのためなどに、適当なダミーデータを簡単に作成することができます。

このライブラリはPHPの[Faker](https://github.com/fzaninotto/Faker)を、修正と変更を加え、Goへ移植したものになります。

## Requirements

```
Go >= 1.22
```

## Installation

```
go get github.com/yrichika/gfaker/pkg/fk
```

## gfakerのインスタンス作成

```go
// デフォルトでは、localeがen_USで作成されます
f := fk.Create()

// 日本語のロケールで作成する場合は、CreateWithLocale()を使います。
j := ja_JP.New()
jp := fk.CreateWithLocale(j)
```

## メソッド


### Rand プリミティブ型のフェイク

`Rand`の

#### Bool

```go
f := fk.Create()

// true/falseが50%ずつの確率
fake := f.Rand.Bool.Evenly() // example: true

// trueが80%, falseが20の確率で返ります
fake := f.Rand.Bool.WeightedToTrue(0.8) // example: true
```

#### Num

```go
// 1から10までのIntを返します。引数に渡した数字が含まれた、ランダムなIntです。
// 例えば、ここでは、1と10は、ランダムな値に含まれます。
f.Rand.Num.IntBt(1, 10) // example: 5

f.Rand.Num.Int32Bt(1, 10) // example: 4

f.Rand.Num.Int64Bt(1, 10) // example: 3

f.Rand.Num.Float32Bt(1.0, 10.0) // example: +2.056392e+000

f.Rand.Num.Float64Bt(1.0, 10.0) // example: +3.627652e+000

// rand.Randのメソッドを使いたい場合は、エイリアスが用意されています
f.Rand.Num.Int()
f.Rand.Num.Intn(10)
```

#### Str 

```go
// アルファベット、数字、特殊文字を含むランダムな文字
f.Rand.Str.Char() // example: "y"
// アルファベット1文字
f.Rand.Str.Letter() // example: "a"
// 数字1文字
f.Rand.Str.Digit() // example: "1"
f.Rand.Str.AlphaRange(5, 10) // example: "VLkwXtKTJ""
f.Rand.Str.AlphaFixedLength(10) // example: "PQRpBVWHow"
// #は数字に、?はアルファベットに、*の英数のどちらかに置き換わります
f.Rand.Str.AlphaDigitsLike("###-???-***") // example: "391-lwe-11u"

```

#### Time

```go
past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)

f.Rand.Time.PastFuture() // example: 2022-01-08 12:24:06.622832978 +0900 JST

f.Rand.Time.PastFrom(past30Years) // example: 1995-05-25 20:49:02.288568665 +0900 JST m=-912941844.443213042

f.Rand.Time.Past() // example: 2002-10-09 19:18:57.246421312 +0900 JST m=-680185687.650174271

f.Rand.Time.FutureTo(future30Years) // example: 2045-03-13 10:18:07.04601675 +0900 JST m=+658635502.851367918

f.Rand.Time.Future() // example: 2041-07-25 21:11:12.119502619 +0900 JST m=+544021834.272691954

f.Rand.Time.TimeRange(past30Years, future30Years) // example: 2016-02-24 08:21:32.32422701 +0900 JST m=-258075587.937158531

f.Rand.Time.Duration() // example: 309679h21m56.609248762s

f.Rand.Time.DurationMilliSec() // example: 875.892572ms

f.Rand.Time.DurationMin() // example: 46m29.429733821s

f.Rand.Time.DurationHour() // example: 8h19m52.645864323s

f.Rand.Time.DurationTo(1 * time.Second) // example: 798.391093ms

f.Rand.Time.DurationRange(1*time.Second, 2*time.Second) // example: 1.818209421s


```

#### Slice

```go

f.Rand.Slice.IntElem([]int{1, 2, 3}) // example: 2

f.Rand.Slice.StrElem([]string{"foo", "bar", "bazz"}) // example: "foo"

```

#### Map

```go
simpleValues := map[any]any{
	"key1": "value1",
  "key2": "value2",
  "key3": "value3",
  "key4": "value4",
}

f.Rand.Map.KeyValue(simpleValues) // example: key4, value4

sliceValues := map[any][]any{
		1: {"value11", "value12"},
		2: {"value21", "value22"},
}

f.Rand.Map.KeySliceValue(sliceValues) // example: 1, [value11 value12]

```


## Global

ロケールに関係なく、同じデータが作成されます


### Barcode

```go
f.Barcode.Ean8() // example: "58594605"

f.Barcode.Ean13() // example: 5945059001019

f.Barcode.Isbn10() // example: 4509472889

f.Barcode.Isbn13() // example: 9787672549372

```

### Color

```go
f.Color.SafeName() // example: 

f.Color.Name() // example: 

f.Color.Hex() // example: 

f.Color.SafeHex() // example: 

f.Color.RgbAsNum() // example: 

f.Color.RgbAsStr() // example: 

f.Color.RgbAsArr() // example: 

f.Color.RgbCss() // example: 

f.Color.RgbaCss() // example: 

f.Color.HslAsNum() // example: 

f.Color.HslAsStr() // example: 

f.Color.HslAsArr() // example: 
```

### File

```go
f.File.MimeType() // example: 

f.File.Extension() // example: 

destDir := "./tmp"
content := "Hello, World!"
returnFullPath := false
f.File.WriteWithText(destDir, content, "txt", returnFullPath)

srcFilePath := "./file/sample.txt"
f.File.CopyFrom(destDir, srcFilePath, "txt", returnFullPath)
```

### Image

```go
// バイナリのイメージデータ
binary, err := f.Image.Binary(100, 100, image.JPG)

// base64にエンコードされたイメージのstring
bs64Str, err := f.Image.Base64(100, 100, image.JPG)

// Goでそのまま扱える`image.Image`のオブジェクトを返します
obj, err := f.Image.Object(100, 100, image.JPG)

```

### Internet

```go
f.Internet.UserName() // example: 

f.Internet.DomainWord() // example: 

f.Internet.Tld() // example: 

f.Internet.DomainName() // example: 

f.Internet.Email() // example: 

f.Internet.Password() // example: 

f.Internet.Ipv4() // example: 

f.Internet.Ipv6() // example: 

f.Internet.LocalIpv4() // example: 

f.Internet.MacAddress() // example: 


```

### Lorem

```go
f.Lorem.Word() // example: 

f.Lorem.WordSliceFixedLength(5) // example: 

f.Lorem.WordSlice(5) // example: 

f.Lorem.Words(5) // example: 

f.Lorem.SentenceFixedLength(5) // example: 

f.Lorem.Sentence(5) // example: 

f.Lorem.SentenceSliceFixedLength(5, 5) // example: 

f.Lorem.SentenceSlice(5, 5) // example: 

f.Lorem.Sentences(5, 5) // example: 

f.Lorem.ParagraphSliceFixedLength(5, 5) // example: 

f.Lorem.ParagraphSlice(5, 5) // example: 

f.Lorem.Paragraphs(5, 5) // example: 


```


### Locale

インスタンス作成時に、渡すロケールによって作成されるデータが変わります。

### Address

```go
f.Address.CitySuffix()

f.Address.CityPrefix()

f.Address.CityName()

f.Address.City()

f.Address.StreetSuffix()

f.Address.StreetName()

f.Address.Street()

f.Address.BuildingNumber()

f.Address.SecondaryAddress()

f.Address.StreetAddress()

f.Address.Postcode()

f.Address.StateAbbr()

f.Address.State()

f.Address.Address()

f.Address.Country()

f.Address.Prefecture()

f.Address.WardSuffix()

f.Address.WardName()

f.Address.Ward()

f.Address.AreaName()

f.Address.AreaNumber()

f.Address.Area()

f.Address.BuildingName()

f.Address.RoomNumber()

f.Address.Latitude()

f.Address.Longitude()

f.Address.LocalCoordinates()


```

### Company

```go
f.Company.CompanyName()

f.Company.CompanyPrefix()

f.Company.CompanySuffix()

f.Company.Name()

f.Company.JobTitleName()

f.Company.JobTitle()

f.Company.EinPrefix()

f.Company.Ein()

```

### Person


```go
f.Person.FirstNameMale()

f.Person.FirstNameFemale()

f.Person.FirstName()

f.Person.LastName()

f.Person.TitleMale()

f.Person.TitleFemale()

f.Person.Title()

f.Person.Suffix()

f.Person.MaleName()

f.Person.FemaleName()

f.Person.Name()

f.Person.Ssn()

f.Person.FirstKanaNameMale()

f.Person.FirstKanaNameFemale()

f.Person.FirstKanaName()

f.Person.LastKanaName()

f.Person.MaleKanaName()

f.Person.FemaleKanaName()

f.Person.KanaName()


```


---

WORKING: まだこれら以外にもメソッドを追加中です。2024年中には一通り完了する予定です。

- Medical
- Miscellaneous
- Payment
- PhoneNumber
- Text
- UserAgent
