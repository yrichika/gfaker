package global

import (
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
)

func CreateInternets() *provider.Internets {
	return &provider.Internets{
		// username
		FirstNames:      *FirstNames,
		LastNames:       *LastNames,
		UserNameFormats: UserNameFormats,
		CreateUserName:  CreateUserName,
		// email
		EmailFormats: EmailFormats,
		CreateEmail:  CreateEmail,
		//
		FreeEmailDomains: FreeEmailDomains,
		Tld:              Tld,
		UrlFormats:       UrlFormats,
		LocalIpBlocks:    LocalIpBlocks,
		// TODO: 他のlocaleのデータも見て、globalで必要そうなら追加する
	}
}

// UserName
var LastNames = &en_US.LastNames

// TODO: concat with male names
var FirstNames = &en_US.FirstNameFemales
var UserNameFormats = []string{
	"{{.LastName}}.{{.FirstName}}",
	"{{.FirstName}}.{{.LastName}}",
	"{{.FirstName}}##",
	"?{{.LastName}}",
}

type UserName struct {
	FirstName string
	LastName  string
}

type UserNameGenerator interface {
	FirstName() string
	LastName() string
}

func CreateUserName(i any) any {
	g := i.(UserNameGenerator)
	return &UserName{
		g.FirstName(),
		g.LastName(),
	}
}

// Email
var FreeEmailDomains = []string{
	"gmail.com", "yahoo.com", "hotmail.com",
}

var Tld = []string{
	"com", "com", "com", "com", "com", "com", "biz", "info", "net", "org",
}

var EmailFormats = []string{
	"{{.UserName}}@{{.DomainName}}",
	"{{.UserName}}@{{.FreeEmailDomain}}",
}

type Email struct {
	UserName        string
	DomainName      string
	FreeEmailDomain string
}

type EmailGenerator interface {
	UserName() string
	DomainName() string
	FreeEmailDomain() string
}

func CreateEmail(i any) any {
	g := i.(EmailGenerator)
	return &Email{
		UserName:        g.UserName(),
		DomainName:      g.DomainName(),
		FreeEmailDomain: g.FreeEmailDomain(),
	}
}

// WORKING:

var LocalIpBlocks = [][]string{
	{"10.0.0.0", "10.255.255.255"},
	{"172.16.0.0", "172.31.255.255"},
	{"192.168.0.0", "192.168.255.255"},
}

var UrlFormats = []string{
	"http://www.{{.DomainName}}/",
	"http://{{.DomainName}}/",
	"http://www.{{.DomainName}}/{{.Slug}}",
	"http://www.{{.DomainName}}/{{.Slug}}",
	"https://www.{{.DomainName}}/{{.Slug}}",
	"http://www.{{.DomainName}}/{{.Slug}}.html",
	"http://{{.DomainName}}/{{.Slug}}",
	"http://{{.DomainName}}/{{.Slug}}",
	"http://{{.DomainName}}/{{.Slug}}.html",
	"https://{{.DomainName}}/{{.Slug}}.html",
}
