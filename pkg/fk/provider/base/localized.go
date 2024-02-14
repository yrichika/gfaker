package base

type Localized struct {
	// Person, Addressごとに、入れ子の構造体を作る
	Person *Person
	// TODO: 全部のロケールのデータを持つ
}

type Person struct {
	FirstNameMales         []string
	FirstNameFemales       []string
	LastNames              []string
	TitleMales             []string
	TitleFemales           []string
	Suffixes               []string
	FirstKanaNameMales     []string
	FirstKanaNameFemales   []string
	LastKanaNames          []string
	MaleNameFormats        []string
	FemaleNameFormats      []string
	CreatePersonNameMale   func(any) any
	CreatePersonNameFemale func(any) any
}
