package color

import (
	"fmt"

	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
)

type Color struct {
	rand *core.Rand
	data *base.Colors
}

func New(rand *core.Rand, global *base.Global) *Color {
	return &Color{
		rand,
		global.Colors,
	}
}

// example 'blue'
func (c *Color) SafeColorName() string {
	return c.rand.Arr.StrElem(c.data.SafeColors)
}

// example 'NavajoWhite'
func (c *Color) ColorName() string {
	return c.rand.Arr.StrElem(c.data.AllColorNames)
}

// example '#fa3cc2'
func (c *Color) HexColor() string {
	number := c.rand.Num.IntBt(1, 16777215)
	return fmt.Sprintf("#%06x", number)
}

// 	 /**
// 	  * @example '#ff0044'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function safeHexColor()
// 	 {
// 		 $color = str_pad(dechex(self::numberBetween(0, 255)), 3, '0', STR_PAD_LEFT);

// 		 return '#' . $color[0] . $color[0] . $color[1] . $color[1] . $color[2] . $color[2];
// 	 }

// 	 /**
// 	  * @example 'array(0,255,122)'
// 	  *
// 	  * @return array
// 	  */
// 	 public static function rgbColorAsArray()
// 	 {
// 		 $color = static::hexColor();

// 		 return [
// 			 hexdec(substr($color, 1, 2)),
// 			 hexdec(substr($color, 3, 2)),
// 			 hexdec(substr($color, 5, 2)),
// 		 ];
// 	 }

// 	 /**
// 	  * @example '0,255,122'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function rgbColor()
// 	 {
// 		 return implode(',', static::rgbColorAsArray());
// 	 }

// 	 /**
// 	  * @example 'rgb(0,255,122)'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function rgbCssColor()
// 	 {
// 		 return 'rgb(' . static::rgbColor() . ')';
// 	 }

// 	 /**
// 	  * @example 'rgba(0,255,122,0.8)'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function rgbaCssColor()
// 	 {
// 		 return 'rgba(' . static::rgbColor() . ',' . static::randomFloat(1, 0, 1) . ')';
// 	 }

// 	 /**
// 	  * @example 'blue'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function safeColorName()
// 	 {
// 		 return static::randomElement(static::$safeColorNames);
// 	 }

// 	 /**
// 	  * @example 'NavajoWhite'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function colorName()
// 	 {
// 		 return static::randomElement(static::$allColorNames);
// 	 }

// 	 /**
// 	  * @example '340,50,20'
// 	  *
// 	  * @return string
// 	  */
// 	 public static function hslColor()
// 	 {
// 		 return sprintf(
// 			 '%s,%s,%s',
// 			 self::numberBetween(0, 360),
// 			 self::numberBetween(0, 100),
// 			 self::numberBetween(0, 100),
// 		 );
// 	 }

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
