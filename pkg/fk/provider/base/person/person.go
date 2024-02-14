package person

import (
	"log"

	"github.com/yrichika/gfaker/pkg/fk/common"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
)

type Person struct {
	rand *core.Rand
	data *base.Person
}

func New(rand *core.Rand, local *base.Localized) *Person {
	return &Person{
		rand,
		local.Person,
	}
}

func (p *Person) FirstNameMale() string {
	return p.rand.Arr.StrElem(p.data.FirstNameMale)
}

func (p *Person) FirstNameFemale() string {
	return p.rand.Arr.StrElem(p.data.FirstNameFemale)
}

func (p *Person) FirstName() string {
	if p.rand.Bool.Evenly() {
		return p.rand.Arr.StrElem(p.data.FirstNameMale)
	}
	return p.rand.Arr.StrElem(p.data.FirstNameFemale)
}

func (p *Person) LastName() string {
	return p.rand.Arr.StrElem(p.data.LastName)
}

func (p *Person) TitleMale() string {
	return p.rand.Arr.StrElem(p.data.TitleMale)
}

func (p *Person) TitleFemale() string {
	return p.rand.Arr.StrElem(p.data.TitleFemale)
}

func (p *Person) Title() string {
	if p.rand.Bool.Evenly() {
		return p.rand.Arr.StrElem(p.data.TitleMale)
	}
	return p.rand.Arr.StrElem(p.data.TitleFemale)
}

func (p *Person) Suffix() string {
	return p.rand.Arr.StrElem(p.data.NameSuffixPerson)
}

func (p *Person) FullNameOf(format string, nameData any) string {

	name, err := common.RenderTemplate(format, nameData)
	if err != nil {
		// TODO: エラーメッセージの共通化
		log.Printf("failed to render template: %v", err)
		return ""
	}
	return name
}

func (p *Person) MaleName() string {
	format := p.rand.Arr.StrElem(p.data.MaleNameFormat)
	nameData := p.data.CreatePersonNameMale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) FemaleName() string {
	format := p.rand.Arr.StrElem(p.data.FemaleNameFormat)
	nameData := p.data.CreatePersonNameFemale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) Name() string {
	if p.rand.Bool.Evenly() {
		return p.MaleName()
	}
	return p.FemaleName()
}

func (p *Person) Ssn() string {
	return p.rand.Str.AlphaDigitsLike("###-##-####")
}

// TODO: ここに、全てのメソッドを定義する。
// en_USでは、カナの名前はないが、そういった場合は、値がnilかチェックして
// log.Printf() でエラーを出すようにする
// 返す値は、"" か、0 など、型に従って、ゼロ値を返すようにする
