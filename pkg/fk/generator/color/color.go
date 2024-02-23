package color

import (
	"fmt"

	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Color struct {
	rand *core.Rand
	data *provider.Colors
}

func New(rand *core.Rand, global *provider.Global) *Color {
	return &Color{
		rand,
		global.Colors,
	}
}

// TEST:
// example 'blue'
func (c *Color) SafeColorName() string {
	return c.rand.Arr.StrElem(c.data.SafeColorNames)
}

// TEST:
// example 'NavajoWhite'
func (c *Color) ColorName() string {
	return c.rand.Arr.StrElem(c.data.AllColorNames)
}

// example '#fa3cc2'
func (c *Color) HexColor() string {
	number := c.rand.Num.IntBt(1, 16777215)
	return fmt.Sprintf("#%06x", number)
}

// example '#ff0044'
func (c *Color) SafeHexColor() string {
	number := c.rand.Num.Intn(256)
	color := fmt.Sprintf("%03x", number)
	return "#" + string(color[0]) + string(color[0]) + string(color[1]) + string(color[1]) + string(color[2]) + string(color[2])
}

// TEST:
// example 0, 255, 122
func (c *Color) RgbColorAsNum() (int, int, int) {
	return c.rand.Num.IntBt(0, 255),
		c.rand.Num.IntBt(0, 255),
		c.rand.Num.IntBt(0, 255)
}

// TEST:
// example '0,255,122'
func (c *Color) RgbColorAsStr() string {
	r, g, b := c.RgbColorAsNum()
	return fmt.Sprintf("%d,%d,%d", r, g, b)
}

// TEST:
// example [0, 255, 122]
func (c *Color) RgbColorAsArr() ([]int, error) {
	r, g, b := c.RgbColorAsNum()
	return []int{r, g, b}, nil

}

// TEST:
// example 'rgb(0,255,122)'
func (c *Color) RgbCssColor() string {
	return "rgb(" + c.RgbColorAsStr() + ")"
}

// example 'rgba(0,255,122,0.8)'
func (c *Color) RgbaCssColor() string {
	return "rgba(" + c.RgbColorAsStr() + "," + fmt.Sprintf("%.1f", c.rand.Num.Float32Bt(0, 1)) + ")"
}

//	 /**
//	  * @example '340,50,20'
//	  *
//	  * @return string
//	  */
//	 public static function hslColor()
//	 {
//		 return sprintf(
//			 '%s,%s,%s',
//			 self::numberBetween(0, 360),
//			 self::numberBetween(0, 100),
//			 self::numberBetween(0, 100),
//		 );
//	 }
//

// TEST:
// example 340, 50, 20
func (c *Color) HslColorAsNum() (int, int, int) {
	return c.rand.Num.IntBt(0, 360),
		c.rand.Num.IntBt(0, 100),
		c.rand.Num.IntBt(0, 100)
}

// TEST:
// example '340,50,20'
func (c *Color) HslColorAsStr() string {
	h, s, l := c.HslColorAsNum()
	return fmt.Sprintf("%d,%d,%d", h, s, l)
}

// 	 /**
// 	  * @example array(340, 50, 20)
// 	  *
// 	  * @return array
// 	  */
// 	 public static function hslColorAsArray()
// 	 {
// 		 return [
// 			 self::numberBetween(0, 360),
// 			 self::numberBetween(0, 100),
// 			 self::numberBetween(0, 100),
// 		 ];
// 	 }
//  }

// TEST:
// example [340, 50, 20]
func (c *Color) HslColorAsArr() []int {
	h, s, l := c.HslColorAsNum()
	return []int{h, s, l}
}
