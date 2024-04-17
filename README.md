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

```

#### Time

```go

```

#### Slice

```go

```

#### Map

```go

```


## Global

ロケールに関係なく、同じデータが作成されます


### Barcode

```go

```

### Color

```go

```

### File

```go

```

### Image

```go

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
