package provider

type Global struct {
	Colors    *Colors
	Files     *Files
	Images    *Images
	Internets *Internets
	Lorems    *Lorems
	// NOTICE: All fields name should be PLURAL
}

type Colors struct {
	SafeColorNames []string
	AllColorNames  []string
}

type Files struct {
	// REFACTOR: MimeTypes should be `MimeTypesAndExtensions`?
	// MimeTypes type should be map[string][]string
	// because of type restriction, it is set as map[any][]any
	MimeTypes map[any][]any
}

type Images struct{}

type Internets struct {
	// username
	FirstNames      []string
	LastNames       []string
	UserNameFormats []string
	CreateUserName  func(any) any
	// email
	EmailFormats []string
	Tld          []string
	CreateEmail  func(any) any
	//
	UrlFormats    []string
	LocalIpBlocks [][]string
}

type Lorems struct {
	Words []string
}
