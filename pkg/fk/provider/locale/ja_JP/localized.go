package ja_JP

import (
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/base/person"
)

func New() *base.Localized {
	return &base.Localized{
		FirstNameMale:          firstNameMale,
		FirstNameFemale:        firstNameFemale,
		LastName:               lastName,
		FirstKanaNameMale:      firstKanaNameMale,
		FirstKanaNameFemale:    firstKanaNameFemale,
		LastKanaName:           lastKanaName,
		MaleNameFormat:         maleNameFormat,
		FemaleNameFormat:       femaleNameFormat,
		CreatePersonNameMale:   createJaJpPersonNameMale,
		CreatePersonNameFemale: createJaJpPersonNameFemale,
	}
}

var maleNameFormat = []string{
	"{{.LastName}} {{.FirstName}}",
}

var femaleNameFormat = []string{
	"{{.LastName}} {{.FirstName}}",
}

type JaJpPersonName struct {
	FirstName string
	LastName  string
}

func createJaJpPersonNameMale(p interface{}) any {
	a := p.(*person.Person)
	return JaJpPersonName{
		FirstName: a.FirstNameMale(),
		LastName:  a.LastName(),
	}
}

func createJaJpPersonNameFemale(p interface{}) any {
	a := p.(*person.Person)
	return JaJpPersonName{
		FirstName: a.FirstNameFemale(),
		LastName:  a.LastName(),
	}
}
