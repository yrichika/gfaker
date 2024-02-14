package base

type Localized struct {
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
	// TODO: 全部のロケールのデータを持つ
}
