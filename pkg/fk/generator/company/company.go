package company

import (
	"fmt"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Company struct {
	rand *core.Rand
	data *provider.Companies
}

func New(rand *core.Rand, local *provider.Localized) *Company {
	return &Company{
		rand,
		local.Companies,
	}
}

func (c *Company) CompanyName() string {
	if len(c.data.CompanyNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.CompanyNames)

}

func (c *Company) CompanyPrefix() string {
	if len(c.data.CompanyPrefixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.CompanyPrefixes)
}

func (c *Company) CompanySuffix() string {
	if len(c.data.CompanySuffixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.CompanySuffixes)
}

func (c *Company) Name() string {
	if len(c.data.CompanyFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := c.rand.Slice.StrElem(c.data.CompanyFormats)
	nameData := c.data.CreateCompany(c)
	return util.RenderTemplate(format, nameData)

}

func (c *Company) JobTitleName() string {
	if len(c.data.JobTitleNames) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.JobTitleNames)
}

func (c *Company) JobTitle() string {
	if len(c.data.JobTitleFormats) == 0 {
		log.UnavailableLocale(1)
		return ""
	}

	format := c.rand.Slice.StrElem(c.data.JobTitleFormats)
	nameData := c.data.CreateJobTitle(c)
	return util.RenderTemplate(format, nameData)
}

func (c *Company) EinPrefix() string {
	if len(c.data.EinPrefixes) == 0 {
		log.UnavailableLocale(1)
		return ""
	}
	return c.rand.Slice.StrElem(c.data.EinPrefixes)
}

// Employer Identification Number (EIN)
// See: https://en.wikipedia.org/wiki/Employer_Identification_Number
// example: "12-3456789"
func (c *Company) Ein() string {
	prefix := c.EinPrefix()
	suffixNum := c.rand.Num.Int64Bt(0, 9999999)
	return prefix + "-" + fmt.Sprintf("%07d", suffixNum)
}
