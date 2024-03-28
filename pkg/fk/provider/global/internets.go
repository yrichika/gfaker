package global

import "github.com/yrichika/gfaker/pkg/fk/provider"

func CreateInternets() *provider.Internets {
	return &provider.Internets{
		FreeEmailDomains: FreeEmailDomains,
		Tld:              Tld,
		UserNameFormats:  UserNameFormats,
		EmailFormats:     EmailFormats,
		UrlFormats:       UrlFormats,
		LocalIpBlocks:    LocalIpBlocks,
		// TODO: 他のlocaleのデータも見て、globalで必要そうなら追加する
	}
}

var FreeEmailDomains = []string{
	"gmail.com", "yahoo.com", "hotmail.com",
}

var Tld = []string{
	"com", "com", "com", "com", "com", "com", "biz", "info", "net", "org",
}

var LocalIpBlocks = [][]string{
	{"10.0.0.0", "10.255.255.255"},
	{"172.16.0.0", "172.31.255.255"},
	{"192.168.0.0", "192.168.255.255"},
}

var UserNameFormats = []string{
	"{{LastName}}.{{FirstName}}",
	"{{FirstName}}.{{LastName}}",
	"{{FirstName}}##",
	"?{{LastName}}",
}

var EmailFormats = []string{
	"{{UserName}}@{{DomainName}}",
	"{{UserName}}@{{FreeEmailDomain}}",
}

var UrlFormats = []string{
	"http://www.{{DomainName}}/",
	"http://{{DomainName}}/",
	"http://www.{{DomainName}}/{{Slug}}",
	"http://www.{{DomainName}}/{{Slug}}",
	"https://www.{{DomainName}}/{{Slug}}",
	"http://www.{{DomainName}}/{{Slug}}.html",
	"http://{{DomainName}}/{{Slug}}",
	"http://{{DomainName}}/{{Slug}}",
	"http://{{DomainName}}/{{Slug}}.html",
	"https://{{DomainName}}/{{Slug}}.html",
}
