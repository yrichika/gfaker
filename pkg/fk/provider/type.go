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
	Companies *Companies
	// TODO: 全部のロケールのデータを持つ
}

type People struct {
	FirstNameMales    []string
	FirstNameFemales  []string
	LastNames         []string
	TitleMales        []string
	TitleFemales      []string
	Suffixes          []string
	MaleNameFormats   []string
	CreateNameMale    func(any) any
	FemaleNameFormats []string
	CreateNameFemale  func(any) any
	// ja_JP
	FirstKanaNameMales   []string
	FirstKanaNameFemales []string
	LastKanaNames        []string
	CreateKanaNameMale   func(any) any
	CreateKanaNameFemale func(any) any
}

type Addresses struct {
	Countries   []string
	Postcodes   []string
	States      []string
	StateAbbrs  []string
	Prefectures []string
	// city
	CityNames    []string
	CityPrefixes []string
	CitySuffixes []string
	CityFormats  []string
	CreateCity   func(any) any
	// ward
	WardNames    []string
	WardSuffixes []string
	WardFormats  []string
	CreateWard   func(any) any
	// area
	AreaNames   []string
	AreaNumbers []string
	AreaFormats []string
	CreateArea  func(any) any
	// street
	StreetNames    []string
	StreetSuffixes []string
	StreetFormats  []string
	CreateStreet   func(any) any
	// street address
	StreetAddressFormats []string
	CreateStreetAddress  func(any) any
	// secondary address
	BuildingNames           []string
	BuildingNumbers         []string
	RoomNumbers             []string
	SecondaryAddressFormats []string
	CreateSecondaryAddress  func(any) any

	AddressFormats []string
	CreateAddress  func(any) any
}

type Companies struct {
	// company name
	CompanyNames    []string
	CompanyPrefixes []string
	CompanySuffixes []string
	CompanyFormats  []string
	CreateCompany   func(any) any
	// job title
	JobTitleNames   []string
	JobTitleFormats []string
	CreateJobTitle  func(any) any
	// EIN
	EinPrefixes []string
}
