package internet

import (
	"strings"

	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Internet struct {
	rand *core.Rand
	data *provider.Internets
}

func New(rand *core.Rand, global *provider.Global) *Internet {
	return &Internet{
		rand,
		global.Internets,
	}
}

func (i *Internet) FirstName() string {
	fName := i.rand.Slice.StrElem(i.data.FirstNames)
	without1Quote := strings.ReplaceAll(fName, "'", "")
	return strings.ToLower(without1Quote)
}

func (i *Internet) LastName() string {
	lName := i.rand.Slice.StrElem(i.data.LastNames)
	without1Quote := strings.ReplaceAll(lName, "'", "")
	return strings.ToLower(without1Quote)
}

func (i *Internet) UserName() string {
	baseFormat := i.rand.Slice.StrElem(i.data.UserNameFormats)
	format := i.rand.Str.AlphaDigitsLike(baseFormat)
	userName := i.data.CreateUserName(i)
	return util.RenderTemplate(format, userName)
}

func (i *Internet) FreeEmailDomain() string {
	return i.rand.Slice.StrElem(i.data.FreeEmailDomains)
}

func (i *Internet) DomainWord() string {
	lastName := i.rand.Slice.StrElem(i.data.LastNames)
	word := strings.ToLower(lastName)
	return word
}

func (i *Internet) Tld() string {
	return i.rand.Slice.StrElem(i.data.Tld)
}

// example: "howell.com"
func (i *Internet) DomainName() string {
	return i.DomainWord() + "." + i.Tld()
}

// example: "jude.borer@oberbrunner.com"
func (i *Internet) Email() string {
	format := i.rand.Slice.StrElem(i.data.EmailFormats)
	data := i.data.CreateEmail(i)
	return util.RenderTemplate(format, data)
}

func (i *Internet) SafeEmail() string {
	//
	return ""
}

func (i *Internet) FreeEmail() string {
	//
	return ""
}

func (i *Internet) CompanyEmail() string {
	//
	return ""
}

func (i *Internet) SafeEmailDomain() string {
	//
	return ""
}

// example: "18w50q2412G5Iky60QL"
func (i *Internet) Password() string {
	min := 8
	max := 20
	length := i.rand.Num.IntBt(min, max)
	like := strings.Repeat("*", length)
	return i.rand.Str.AlphaDigitsLike(like)
}

func (i *Internet) Url() string {
	//
	return ""
}

func (i *Internet) Slug() string {
	//
	return ""
}

func (i *Internet) Ipv4() string {
	//
	return ""
}

func (i *Internet) Ipv6() string {
	//
	return ""
}

func (i *Internet) LocalIpv4() string {
	//
	return ""
}

func (i *Internet) MacAddress() string {
	//
	return ""
}
