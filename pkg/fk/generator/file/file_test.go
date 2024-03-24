package file

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
)

func TestFile(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Files: global.CreateFiles(),
	}

	file := New(coreRand, global)

	t := gt.CreateTest(testingT)
	t.Describe("File", func() {
		t.Test("MimeType should return a string", func() {
			mimeType := file.MimeType()

			_, exists := global.Files.MimeTypes[mimeType]
			gt.Expect(t, &exists).ToBe(true)
		})

		t.Test("Extension should return a string", func() {
			extension := file.Extension()

			gt.Expect(t, &extension).ToMatchRegex(`^[0-9a-z-]{3,10}$`)
		})
	})
}
