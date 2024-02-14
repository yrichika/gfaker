package base

type Localized struct {
	// Person, Addressごとに、入れ子の構造体を作る
	Person *Person
	// TODO: 全部のロケールのデータを持つ
}

type Person struct {
	FirstNameMale          []string
	FirstNameFemale        []string
	LastName               []string
	TitleMale              []string
	TitleFemale            []string
	NameSuffixPerson       []string
	FirstKanaNameMale      []string
	FirstKanaNameFemale    []string
	LastKanaName           []string
	MaleNameFormat         []string
	FemaleNameFormat       []string
	CreatePersonNameMale   func(any) any
	CreatePersonNameFemale func(any) any
}
