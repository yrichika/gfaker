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
f.Color.SafeName() // example: "silver"

f.Color.Name() // example: "Gainsboro"

f.Color.Hex() // example: "#0e457a"

f.Color.SafeHex() // example: "#00dd22"

f.Color.RgbAsNum() // example: 15, 247, 177

f.Color.RgbAsStr() // example: "161,181,228"

f.Color.RgbAsArr() // example: [98 35 65]

f.Color.RgbCss() // example: "rgb(223,67,224)"

f.Color.RgbaCss() // example: "rgba(66,112,144,0.3)"

f.Color.HslAsNum() // example: 153, 97, 56

f.Color.HslAsStr() // example: "149,85,59"

f.Color.HslAsArr() // example: [31 69 46]
```

### File

```go
f.File.MimeType() // example: "application/widget"

f.File.Extension() // example: "tga"

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
f.Internet.UserName() // example: ayla.prosacco

f.Internet.DomainWord() // example: klein

f.Internet.Tld() // example: com

f.Internet.DomainName() // example: gutkowski.info

f.Internet.Email() // example: charity.ziemann@mertz.net

f.Internet.Password() // example: 9AYD57IF580tD76p55y

f.Internet.Ipv4() // example: 190.238.68.2

f.Internet.Ipv6() // example: 112a:792e:884c:99e5:d7a0:2b2c:df2b:9c48

f.Internet.LocalIpv4() // example: 172.27.204.249

f.Internet.MacAddress() // example: 01:ED:77:9F:1C:E1


```

### Lorem

```go
f.Lorem.Word() // example: "qui"
// 指定した個数の文字列の配列を返す

f.Lorem.WordSliceFixedLength(5) // example: [consectetur quia reprehenderit est consectetur]

// 指定した文字数を上限としてランダムな個数の文字列のスライスを返す
f.Lorem.WordSlice(5) // example: [nisi porro]

// 指定した文字数を上限としてランダムな個数の文字列を返す
f.Lorem.Words(5) // example: "sint officia eveniet ut sint"

f.Lorem.SentenceFixedLength(5) // example: "Aut corrupti ullam delectus exercitationem."

f.Lorem.Sentence(5) // example: "Vel."

f.Lorem.SentenceSliceFixedLength(5, 5) // example: ["Quis." "Vitae et quisquam." "Earum." "Omnis." "Perferendis eius fugit voluptas qui."]

f.Lorem.SentenceSlice(5, 5) // example: ["Facilis." "Consequatur sed saepe necessitatibus et."]

f.Lorem.Sentences(5, 5) // example: "Voluptatem sed omnis vel repudiandae. Quo et. Sit optio ipsa beatae. Veritatis iusto."

f.Lorem.ParagraphSliceFixedLength(5, 5) // example: 
// ["Et incidunt quia necessitatibus." "Porro ut ipsa nulla quos et dignissimos in." "Voluptas illo consectetur." "Illo." "Doloremque."" Eos placeat." "Nam nostrum sed necessitatibus voluptas provident est quibusdam saepe reprehenderit ut illum quae consequatur excepturi corporis illo voluptatum sint omnis magni qui adipisci voluptatem." "Sed vel veritatis dolores et voluptatum molestiae sequi aut." "Fugiat est ducimus et eos eligendi." "Omnis molestias dolorem animi sapiente voluptatem soluta nostrum qui reprehenderit." "Enim dolor aliquam mollitia beatae omnis autem sunt perspiciatis corrupti molestiae sunt sed qui id facilis laudantium ut eveniet." "Exercitationem quibusdam corporis alias porro vel." "Aut dolorum magni." "Voluptatem libero ipsa." "Eaque."]

f.Lorem.ParagraphSlice(5, 5) // example: 
// ["Neque impedit inventore qui repellendus dolores nulla minima nulla ratione similique illum non asperiores error iusto." "Voluptatem soluta adipisci qui odio magnam fuga consequatur pariatur veniam aut quis ipsam quibusdam voluptatibus et sapiente." "Molestiae quaerat consectetur pariatur possimus." "Nisi et eum quia suscipit itaque magnam architecto porro ut earum vel possimus at commodi aliquid possimus est magni sit et molestias odit animi velit eos numquam animi voluptatum voluptatem." "Deserunt error id consequatur." "A quis." "Eius labore molestiae ut omnis." "Et nam." "Qui nobis aut. ""Consectetur eum aut non dolorem enim voluptas vitae." "Alias ullam voluptas est voluptatem dolore reprehenderit." "Voluptate qui possimus animi ut voluptatem quo asperiores." "Veritatis voluptatibus."]

f.Lorem.Paragraphs(5, 5) // example: 
// "Veniam sed enim quidem blanditiis excepturi dicta molestias numquam enim. Nisi ipsum reiciendis vel voluptatum dolorum eum deleniti voluptas eum sed rem nulla."

```


### Locale

インスタンス作成時に、渡すロケールによって作成されるデータが変わります。

### Address

```go
f.Address.CitySuffix() // example: "land"
// example ja_Jp: "村"

f.Address.CityPrefix() // example: "Port"

f.Address.CityName() // example: "Crona"
// example ja_Jp: "日高"

f.Address.City() // example: "Fayside"
// example ja_Jp: "天童市"

f.Address.StreetSuffix() // example: "Parkways"


f.Address.StreetName() // example: "Mraz"

f.Address.Street() // example: "Considine Island"

f.Address.BuildingNumber() // example: "3114"

f.Address.SecondaryAddress() // example: "Suite 760"
// example ja_Jp: "レジデンス加納8K号"

f.Address.StreetAddress() // example: "618 Lynch Apt. 42"

f.Address.Postcode() // example: "75602"
// example ja_Jp: "6858030"

f.Address.StateAbbr() // example: "CA"

f.Address.State() // example: "Colorado"

f.Address.Address() // example: 
// "0730 Gleason Apt. 34\n
// South West, NM 08956"
// example ja_Jp: "334-7397  茨城県大仙町東区東夷川3-7-0"


f.Address.Country() // example: "United States Minor Outlying Islands"
// example ja_Jp: "サウジアラビア"

f.Address.Prefecture() // example: "佐賀県"

f.Address.WardSuffix() // example: "区"

f.Address.WardName() // example: "東"

f.Address.Ward() // example: "北区"

f.Address.AreaName() // example: "北小路"

f.Address.AreaNumber() // example: "2-8-04"

f.Address.Area() // example: "谷口町7丁目1番地610"

f.Address.BuildingName() // example: "笹田"

f.Address.RoomNumber() // example: "804"

f.Address.Latitude() // example: 54.617171

f.Address.Longitude() // example: 74.851822

f.Address.LocalCoordinates() // example: 82.130718, -121.140770


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
