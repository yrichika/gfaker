package global

import "github.com/yrichika/gfaker/pkg/fk/provider/base"

func CreateColors() *base.Colors {
	return &base.Colors{
		SafeColors:    safeColors,
		AllColorNames: allColorNames,
	}
}
