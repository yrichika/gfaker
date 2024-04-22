# gfaker



## 

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
fake := f.Rand.Bool.Evenly()

// trueが80%, falseが20の確率で返ります
fake := f.Rand.Bool.WeightedToTrue(0.8)
```

#### Num

```go
// 1から10までのIntを返します。引数に渡した数字が含まれた、ランダムなIntです。
// 例えば、ここでは、1と10は、ランダムな値に含まれます。
f.Rand.Num.IntBt(1, 10)
f.Rand.Num.Int32Bt(1, 10)
f.Rand.Num.Int64Bt(1, 10)
f.Rand.Num.Float32Bt(1.0, 10.0)
f.Rand.Num.Float64Bt(1.0, 10.0)

// rand.Randのメソッドを使いたい場合は、エイリアスが用意されています
f.Rand.Num.Int()
f.Rand.Num.Intn(10)
```

#### Str 

```go
// アルファベット、数字、特殊文字を含むランダムな文字
f.Rand.Str.Char()
// アルファベット1文字
f.Rand.Str.Letter()
// 数字1文字
f.Rand.Str.Digit()

f.Rand.Str.AlphaRange(5, 10)

f.Rand.Str.AlphaFixedLength(10)
f.Rand.Str.AlphaDigitsLike("###-???-***")

```

#### Time

```go
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


```

#### Slice

```go

f.Rand.Slice.IntElem([]int{1, 2, 3})

f.Rand.Slice.StrElem([]string{"foo", "bar", "bazz"})

```

#### Map

```go
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

f.Rand.Map.KeySliceValue(arrayedValues)

```


## Global

ロケールに関係なく、同じデータが作成されます


### Barcode

```go
f.Barcode.Ean8()

f.Barcode.Ean13()

f.Barcode.Isbn10()

f.Barcode.Isbn13()

```

### Color

```go
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
```

### File

```go
f.File.MimeType()

f.File.Extension()

destDir := "./tmp"
content := "Hello, World!"
returnFullPath := false
f.File.WriteWithText(destDir, content, "txt", returnFullPath)

srcFilePath := "./file/sample.txt"
f.File.CopyFrom(destDir, srcFilePath, "txt", returnFillPath)
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

```

### Lorem

```go

```


### Locale

インスタンス作成時に、渡すロケールによって作成されるデータが変わります。

### Address

```go

```

### Company

```go

```

### Person


```go

```


---

WORKING: まだこれら以外にもメソッドを追加中です。2024年中には一通り完了する予定です。

- Medical
- Miscellaneous
- Payment
- PhoneNumber
- Text
- UserAgent
