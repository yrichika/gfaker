package image

import (
	"bytes"
	"encoding/base64"
	stdimage "image"
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
)

func TestImage(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Images: global.CreateImages(),
	}

	image := New(coreRand, global)

	tBin := gt.CreateTest(testingT)
	tBin.Describe("Binary", func() {
		tBin.Test("should return specified image binary data", func() {
			r, err := image.Binary(100, 100, JPG)
			img, format, _ := stdimage.Decode(bytes.NewReader(r))

			gt.Expect(tBin, &err).Not().ToContainError()
			gt.Expect(tBin, &format).ToBe("jpeg")

			bounds := img.Bounds()
			width := bounds.Dx()
			gt.Expect(tBin, &width).ToBe(100)
			height := bounds.Dy()
			gt.Expect(tBin, &height).ToBe(100)
		})

		tBin.Test("should return 3840px width image even if more than that is specified", func() {
			r, _ := image.Binary(4000, 100, JPG)
			img, _, _ := stdimage.Decode(bytes.NewReader(r))

			bounds := img.Bounds()
			width := bounds.Dx()
			gt.Expect(tBin, &width).ToBe(maxHeightWidth)
		})

		tBin.Test("should return 3840px height image even if more than that is specified", func() {
			r, _ := image.Binary(100, 4000, JPG)
			img, _, _ := stdimage.Decode(bytes.NewReader(r))

			bounds := img.Bounds()
			height := bounds.Dy()
			gt.Expect(tBin, &height).ToBe(maxHeightWidth)
		})

		tBin.Test("should return default jpg image even if other format specified", func() {
			r, _ := image.Binary(100, 100, "other_image_format")
			_, format, _ := stdimage.Decode(bytes.NewReader(r))

			gt.Expect(tBin, &format).ToBe("jpeg")
		})
	})

	tObj := gt.CreateTest(testingT)
	tObj.Describe("Object", func() {
		tObj.Test("should return specified image object", func() {
			img, err := image.Object(100, 100, JPG)

			gt.Expect(tObj, &err).Not().ToContainError()

			bounds := img.Bounds()
			width := bounds.Dx()
			gt.Expect(tObj, &width).ToBe(100)
			height := bounds.Dy()
			gt.Expect(tObj, &height).ToBe(100)
		})

		tObj.Test("log output when width is more than max width", func() {
			img, _ := image.Object(4000, 100, JPG)

			bounds := img.Bounds()
			width := bounds.Dx()
			gt.Expect(tObj, &width).ToBe(maxHeightWidth)
		})
	})

	t64 := gt.CreateTest(testingT)
	t64.Describe("Base64", func() {
		t64.Test("should return base64 string image", func() {
			r, err := image.Base64(100, 100, JPG)
			img, _ := base64ToImage(r)

			gt.Expect(t64, &err).Not().ToContainError()

			bounds := img.Bounds()
			width := bounds.Dx()
			gt.Expect(t64, &width).ToBe(100)
			height := bounds.Dy()
			gt.Expect(t64, &height).ToBe(100)
		})

		t64.Test("log output when width is more than max width", func() {
			r, _ := image.Base64(4000, 100, JPG)
			img, _ := base64ToImage(r)

			bounds := img.Bounds()
			width := bounds.Dx()
			gt.Expect(t64, &width).ToBe(maxHeightWidth)
		})
	})
}

func base64ToImage(base64Str string) (stdimage.Image, error) {
	// Decode the base64 string to binary data
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, err
	}

	img, _, err := stdimage.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	return img, nil
}
