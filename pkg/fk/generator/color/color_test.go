package color

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
)

func TestColor(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Colors: global.CreateColors(),
	}

	color := New(coreRand, global)

	tScn := gt.CreateTest(testingT)
	tScn.Describe("SafeColorName", func() {
		tScn.Todo("")
	})

	tCn := gt.CreateTest(testingT)
	tCn.Describe("ColorName", func() {
		tCn.Todo("")
	})

	tHc := gt.CreateTest(testingT)
	tHc.Describe("HexColor", func() {
		tHc.It("should return random hex color string", func() {
			r := color.HexColor()
			gt.Expect(tHc, &r).ToMatchRegex(`^#[0-9a-f]{6}$`)
		})

	})

	tShc := gt.CreateTest(testingT)
	tShc.Describe("SafeHexColor", func() {
		tShc.It("should return random safe hex color string", func() {
			r := color.SafeHexColor()
			// FIXME: this regex is not exactly right.
			gt.Expect(tShc, &r).ToMatchRegex(`^#[0-9a-f]{6}$`)
		})
	})

	// TODO: 他のテスト

	tRgbaCss := gt.CreateTest(testingT)
	tRgbaCss.Describe("RgbaCssColor", func() {
		tRgbaCss.It("should return random rgba css color string", func() {
			r := color.RgbaCssColor()
			gt.Expect(tRgbaCss, &r).ToMatchRegex(`^rgba\(\d{1,3},\d{1,3},\d{1,3},[01]\.\d{1}\)$`)
		})
	})
}
