package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/gif"
	"image/jpeg"
	"image/png"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Image struct {
	rand    *core.Rand
	data    *provider.Images
	logSkip int
}

func New(rand *core.Rand, global *provider.Global) *Image {
	return &Image{
		rand,
		global.Images,
		1,
	}
}

type ImageFormat string

const (
	JPG = ImageFormat("jpg")
	PNG = ImageFormat("png")
	GIF = ImageFormat("gif")
)

const maxHeightWidth = 3840

// the image content will be just white blank image
// width and height should be less than or eq 3840px
func (i *Image) Binary(width int, height int, format ImageFormat) ([]byte, error) {
	imgWidth := width
	if width > maxHeightWidth {
		imgWidth = maxHeightWidth
		log.GeneralError("Image width is too large, it will be set to 3840px", i.logSkip)
	}
	imgHeight := height
	if height > maxHeightWidth {
		imgHeight = maxHeightWidth
		log.GeneralError("Image height is too large, it will be set to 3840px", i.logSkip)
	}

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Set each pixel color to white
	for y := 0; y < imgHeight; y++ {
		for x := 0; x < imgWidth; x++ {
			img.Set(x, y, color.RGBA{255, 255, 255, 255})
		}
	}

	buf := new(bytes.Buffer)
	var err error
	switch format {
	case JPG:
		err = jpeg.Encode(buf, img, nil)
	case PNG:
		err = png.Encode(buf, img)
	case GIF:
		err = gif.Encode(buf, img, nil)
	default:
		// default to jpg
		err = jpeg.Encode(buf, img, nil)
		log.GeneralError("Image format must be either [image.JPG], [image.PNG] or [image.GIF]. If not any of those, it defaults to JPG.", 1)
	}

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// the image content will be just white blank image
// width and height should be less than or eq 3840px
func (i *Image) Object(width int, height int, format ImageFormat) (image.Image, error) {
	i.logSkip = 2
	binary, err := i.Binary(width, height, format)
	i.logSkip = 1 // reset logSkip
	if err != nil {
		return nil, err
	}

	obj, _, errDecode := image.Decode(bytes.NewReader(binary))
	return obj, errDecode
}

// the image content will be just white blank image
// width and height should be less than or eq 3840px
func (i *Image) Base64(width int, height int, format ImageFormat) (string, error) {
	i.logSkip = 2
	img, err := i.Binary(width, height, format)
	i.logSkip = 1 // reset logSkip

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(img), nil
}
