package person

import (
	"log"

	"github.com/yrichika/gfaker/pkg/fk/common"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
)

type Person struct {
	rand *core.Rand
	data *base.People
}

func New(rand *core.Rand, local *base.Localized) *Person {
	return &Person{
		rand,
		local.People,
	}
}

func (p *Person) FirstNameMale() string {
	return p.rand.Arr.StrElem(p.data.FirstNameMales)
}

func (p *Person) FirstNameFemale() string {
	return p.rand.Arr.StrElem(p.data.FirstNameFemales)
}

func (p *Person) FirstName() string {
	if p.rand.Bool.Evenly() {
		return p.rand.Arr.StrElem(p.data.FirstNameMales)
	}
	return p.rand.Arr.StrElem(p.data.FirstNameFemales)
}

func (p *Person) LastName() string {
	return p.rand.Arr.StrElem(p.data.LastNames)
}

func (p *Person) TitleMale() string {
	return p.rand.Arr.StrElem(p.data.TitleMales)
}

func (p *Person) TitleFemale() string {
	return p.rand.Arr.StrElem(p.data.TitleFemales)
}

func (p *Person) Title() string {
	if p.rand.Bool.Evenly() {
		return p.rand.Arr.StrElem(p.data.TitleMales)
	}
	return p.rand.Arr.StrElem(p.data.TitleFemales)
}

func (p *Person) Suffix() string {
	return p.rand.Arr.StrElem(p.data.Suffixes)
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
	format := p.rand.Arr.StrElem(p.data.MaleNameFormats)
	nameData := p.data.CreateNameMale(p)
	return p.FullNameOf(format, nameData)
}

func (p *Person) FemaleName() string {
	format := p.rand.Arr.StrElem(p.data.FemaleNameFormats)
	nameData := p.data.CreateNameFemale(p)
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
