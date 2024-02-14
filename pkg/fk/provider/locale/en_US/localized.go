package en_US

import (
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/base/person"
)

func New() *base.Localized {
	return &base.Localized{
		FirstNameMale:          firstNameMale,
		FirstNameFemale:        firstNameFemale,
		LastName:               lastName,
		TitleMale:              titleMale,
		TitleFemale:            titleFemale,
		NameSuffixPerson:       nameSuffixPerson,
		MaleNameFormat:         maleNameFormat,
		FemaleNameFormat:       femaleNameFormat,
		CreatePersonNameMale:   createUsPersonNameMale,
		CreatePersonNameFemale: createUsPersonNameFeMale,
	}
}

var maleNameFormat = []string{
	"{{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}}",
	"{{.Title}} {{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}} {{.Suffix}}",
	"{{.Title}} {{.FirstName}} {{.LastName}} {{.Suffix}}",
}

var femaleNameFormat = []string{
	"{{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}}",
	"{{.Title}} {{.FirstName}} {{.LastName}}",
	"{{.FirstName}} {{.LastName}} {{.Suffix}}",
	"{{.Title}} {{.FirstName}} {{.LastName}} {{.Suffix}}",
}

type UsPersonName struct {
	FirstName string
	LastName  string
	Title     string
	Suffix    string
}

func createUsPersonNameMale(p any) any {
	a := p.(*person.Person)
	return UsPersonName{
		FirstName: a.FirstNameMale(),
		LastName:  a.LastName(),
		Title:     a.TitleMale(),
		Suffix:    a.Suffix(),
	}
}

func createUsPersonNameFeMale(p any) any {
	a := p.(*person.Person)
	return UsPersonName{
		FirstName: a.FirstNameFemale(),
		LastName:  a.LastName(),
		Title:     a.TitleFemale(),
		Suffix:    a.Suffix(),
	}
}
