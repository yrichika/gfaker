package person

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/ja_JP"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestPerson(testingT *testing.T) {
	localized := ja_JP.New()
	coreRand := core.NewRand(util.RandSeed())
	person := New(coreRand, localized)

	t := gt.CreateTest(testingT)
	t.Describe("KanaName", func() {
		t.It("should return a kana name", func() {
			r := person.KanaName()
			testutil.Output("Person.KanaName", r)
		})
	})
}
