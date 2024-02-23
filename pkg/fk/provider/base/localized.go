package base

type Global struct {
	Colors *Colors
	// 名前は、それぞれのカテゴリーの'複数形'にすること
}

type Colors struct {
	SafeColorNames []string
	AllColorNames  []string
}

type Localized struct {
	// People, Addressesごとに、入れ子の構造体を作る
	// 名前は、それぞれのカテゴリーの'複数形'にすること
	People *People
	// TODO: 全部のロケールのデータを持つ
}

type People struct {
	FirstNameMales       []string
	FirstNameFemales     []string
	LastNames            []string
	TitleMales           []string
	TitleFemales         []string
	Suffixes             []string
	FirstKanaNameMales   []string
	FirstKanaNameFemales []string
	LastKanaNames        []string
	MaleNameFormats      []string
	FemaleNameFormats    []string
	CreateNameMale       func(any) any
	CreateNameFemale     func(any) any
}
