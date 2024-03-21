package company

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestCompany(testingT *testing.T) {
	localized := en_US.New()
	coreRand := core.NewRand(util.RandSeed())
	company := New(coreRand, localized)

	t1 := gt.CreateTest(testingT)
	t1.Describe("Company", func() {
		t1.Test("CompanyName should return a company name", func() {
			r := company.CompanyName()
			gt.Expect(t1, &r).ToBeIn(en_US.CompanyNames)
		})

		t1.Test("CompanySuffix should return a company suffix", func() {
			r := company.CompanySuffix()
			gt.Expect(t1, &r).ToBeIn(en_US.CompanySuffixes)
		})

		t1.Test("Name should return a company name", func() {
			r := company.Name()
			testutil.Output("Company.Name", r)
		})
	})

	tJt := gt.CreateTest(testingT)
	tJt.Describe("JobTitle", func() {
		tJt.Test("JobTitleName should return a job title name", func() {
			r := company.JobTitleName()
			gt.Expect(tJt, &r).ToBeIn(en_US.JobTitleNames)
		})

		tJt.Test("JobTitle should return a job title", func() {
			r := company.JobTitle()
			testutil.Output("Company.JobTitle", r)
		})
	})

	tEin := gt.CreateTest(testingT)
	tEin.Describe("EIN", func() {
		tEin.Test("EinPrefix should return a EIN prefix", func() {
			r := company.EinPrefix()
			gt.Expect(tEin, &r).ToBeIn(en_US.EinPrefixes)
		})
		tEin.Test("Ein should return a EIN", func() {
			r := company.Ein()
			gt.Expect(tEin, &r).ToMatchRegex(`\d{2}-\d{7}`)
		})
	})

	jaJp := ja_JP.New()
	companyJaJp := New(coreRand, jaJp)
	tJaJp := gt.CreateTest(testingT)
	tJaJp.Describe("ja_JP Company", func() {
		tJaJp.Test("CompanyPrefix should return a company prefix", func() {
			r := companyJaJp.CompanyPrefix()
			gt.Expect(tJaJp, &r).ToBeIn(ja_JP.CompanyPrefixes)
		})

		tJaJp.Test("Name should return a company name", func() {
			r := companyJaJp.Name()
			testutil.Output("Company.Name", r)
		})
	})
}
