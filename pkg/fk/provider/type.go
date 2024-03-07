package provider

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
	People    *People
	Addresses *Addresses
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
	CreateKanaNameMale   func(any) any
	CreateKanaNameFemale func(any) any
}

type Addresses struct {
	CityPrefixes            []string
	CitySuffixes            []string
	CityNames               []string
	BuildingNumbers         []string
	StreetSuffixes          []string
	Postcodes               []string
	States                  []string
	StateAbbrs              []string
	Prefectures             []string
	Countries               []string
	CityFormats             []string
	StreetNameFormats       []string
	StreetAddressFormats    []string
	SecondaryAddressFormats []string
	AddressFormats          []string
	CreateCityFullName      func(any) any
	CreateStreetName        func(any) any
	CreateStreetAddress     func(any) any
	CreateAddress           func(any) any
}
