package color

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/global"
)

func TestColor(testingT *testing.T) {

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	coreRand := core.NewRand(rand)
	localized := &base.Localized{
		Colors: global.CreateColors(),
	}

	t := gt.CreateTest(testingT)
	color := New(coreRand, localized)
	t.Describe("HexColor", func() {
		t.It("should return random hex color string", func() {
			r := color.HexColor()
			gt.Expect(t, &r).ToMatchRegex(`^#[0-9a-f]{6}$`)
		})
	})
}
