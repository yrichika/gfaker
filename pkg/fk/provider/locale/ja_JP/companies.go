package ja_JP

import "github.com/yrichika/gfaker/pkg/fk/provider"

func CreateCompanies() *provider.Companies {
	return &provider.Companies{
		CompanyNames:    CompanyNames,
		CompanyPrefixes: CompanyPrefixes,
		CompanySuffixes: CompanySuffixes,
		CompanyFormats:  CompanyFormats,
		CreateCompany:   CreateJaJpCompany,
	}
}

var CompanyNames = LastNames
var CompanyPrefixes = []string{"株式会社", "有限会社"}
var CompanySuffixes = CompanyPrefixes

var CompanyFormats = []string{
	"{{.CompanyPrefix}} {{.CompanyName}}",
	"{{.CompanyName}} {{.CompanySuffix}}",
}

type JaJpCompany struct {
	CompanyName   string
	CompanyPrefix string
	CompanySuffix string
}

type JaJpCompanyGenerator interface {
	CompanyName() string
	CompanySuffix() string
	CompanyPrefix() string
}

func CreateJaJpCompany(c any) any {
	g := c.(JaJpCompanyGenerator)
	return &JaJpCompany{
		CompanyName:   g.CompanyName(),
		CompanyPrefix: g.CompanyPrefix(),
		CompanySuffix: g.CompanySuffix(),
	}
}
