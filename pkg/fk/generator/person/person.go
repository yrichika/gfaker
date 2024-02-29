package person

import (
	"github.com/yrichika/gfaker/pkg/fk/common/log"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Person struct {
	rand *core.Rand
	data *provider.People
}

func New(rand *core.Rand, local *provider.Localized) *Person {
	return &Person{
		rand,
		local.People,
	}
}

func (p *Person) FirstNameMale() string {
	if len(p.data.FirstNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstNameMales)
}

func (p *Person) FirstNameFemale() string {
	if len(p.data.FirstNameFemales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstNameFemales)
}

func (p *Person) FirstName() string {
	if len(p.data.FirstNameFemales) == 0 || len(p.data.FirstNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.rand.Slice.StrElem(p.data.FirstNameMales)
	}
	return p.rand.Slice.StrElem(p.data.FirstNameFemales)
}

func (p *Person) LastName() string {
	if len(p.data.LastNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	return p.rand.Slice.StrElem(p.data.LastNames)
}

func (p *Person) TitleMale() string {
	if len(p.data.TitleMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.TitleMales)
}

func (p *Person) TitleFemale() string {
	if len(p.data.TitleFemales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.TitleFemales)
}

func (p *Person) Title() string {
	if len(p.data.TitleFemales) == 0 || len(p.data.TitleMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.rand.Slice.StrElem(p.data.TitleMales)
	}
	return p.rand.Slice.StrElem(p.data.TitleFemales)
}

func (p *Person) Suffix() string {
	if len(p.data.Suffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	return p.rand.Slice.StrElem(p.data.Suffixes)
}

func (p *Person) FullNameOf(format string, nameData any) string {
	return util.RenderTemplate(format, nameData)
}

func (p *Person) MaleName() string {
	if len(p.data.MaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.MaleNameFormats)
	nameData := p.data.CreateNameMale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) FemaleName() string {
	if len(p.data.FemaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.FemaleNameFormats)
	nameData := p.data.CreateNameFemale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) Name() string {
	if (len(p.data.MaleNameFormats) == 0) || (len(p.data.FemaleNameFormats) == 0) {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.MaleName()
	}
	return p.FemaleName()
}

func (p *Person) Ssn() string {
	return p.rand.Str.AlphaDigitsLike("###-##-####")
}

func (p *Person) FirstKanaNameMale() string {
	if len(p.data.FirstKanaNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstKanaNameMales)
}

func (p *Person) FirstKanaNameFemale() string {
	if len(p.data.FirstKanaNameFemales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return p.rand.Slice.StrElem(p.data.FirstKanaNameFemales)
}

func (p *Person) FirstKanaName() string {
	if len(p.data.FirstKanaNameFemales) == 0 || len(p.data.FirstKanaNameMales) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.rand.Slice.StrElem(p.data.FirstKanaNameMales)
	}
	return p.rand.Slice.StrElem(p.data.FirstKanaNameFemales)
}

func (p *Person) LastKanaName() string {
	if len(p.data.LastKanaNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	return p.rand.Slice.StrElem(p.data.LastKanaNames)
}

func (p *Person) MaleKanaName() string {
	if len(p.data.MaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.MaleNameFormats)
	nameData := p.data.CreateKanaNameMale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) FemaleKanaName() string {
	if len(p.data.FemaleNameFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := p.rand.Slice.StrElem(p.data.FemaleNameFormats)
	nameData := p.data.CreateKanaNameFemale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) KanaName() string {
	if (len(p.data.MaleNameFormats) == 0) || (len(p.data.FemaleNameFormats) == 0) {
		log.UnavailableLocale(1)
		return ""
	}

	if p.rand.Bool.Evenly() {
		return p.MaleKanaName()
	}
	return p.FemaleKanaName()
}
