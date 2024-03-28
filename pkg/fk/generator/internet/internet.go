package internet

import (
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

func (i *Internet) Email() string {
	//
	return ""
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

func (i *Internet) FreeEmailDomain() string {
	//
	return ""
}

func (i *Internet) SafeEmailDomain() string {
	//
	return ""
}

func (i *Internet) UserName() string {
	//
	return ""
}

func (i *Internet) Password() string {
	//
	return ""
}

func (i *Internet) DomainName() string {
	//
	return ""
}

func (i *Internet) DomainWord() string {
	//
	return ""
}

func (i *Internet) Tld() string {
	//
	return ""
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
