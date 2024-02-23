package global

import "github.com/yrichika/gfaker/pkg/fk/provider/base"

func New() *base.Global {
	return &base.Global{
		Colors: CreateColors(),
	}
}

func CreateColors() *base.Colors {
	return &base.Colors{
		SafeColorNames: safeColorNames,
		AllColorNames:  allColorNames,
	}
}
