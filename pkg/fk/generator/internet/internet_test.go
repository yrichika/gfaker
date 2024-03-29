package internet

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestInternet(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Internets: global.CreateInternets(),
	}

	inet := New(coreRand, global)

	t := gt.CreateTest(testingT)
	t.Describe("Email", func() {
		// TODO: FirstName, LastNameがlowerになってるかをテスト。可能か?

		t.Test("UserName should return a user name", func() {
			r := inet.UserName()
			testutil.Output("Internet.UserName", r)
		})

		t.Test("FreeEmailDomain should return a free email domain", func() {
			r := inet.FreeEmailDomain()
			gt.Expect(t, &r).ToBeIn(inet.data.FreeEmailDomains)
		})

		t.Test("DomainWord should return a domain word", func() {
			r := inet.DomainWord()
			// TODO: アサート方法を考える lowercaseになっているか
			testutil.Output("Internet.DomainWord", r)
		})

		t.Test("Tld should return a tld", func() {
			r := inet.Tld()
			gt.Expect(t, &r).ToBeIn(inet.data.Tld)
		})

		t.Test("DomainName should return a domain name", func() {
			r := inet.DomainName()
			testutil.Output("Internet.DomainName", r)
		})

		t.Test("Email should return an email", func() {
			r := inet.Email()
			testutil.Output("Internet.Email", r)
		})

		// TODO:

		t.Test("Password should return a random string between 8 to 20 length", func() {
			r := inet.Password()
			gt.Expect(t, &r).ToMatchRegex(`^[\d\w]{8,20}$`)
		})
	})
}
