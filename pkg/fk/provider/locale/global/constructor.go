package global

import "github.com/yrichika/gfaker/pkg/fk/provider"

func New() *provider.Global {
	return &provider.Global{
		Colors: CreateColors(),
	}
}

func CreateColors() *provider.Colors {
	return &provider.Colors{
		SafeColorNames: safeColorNames,
		AllColorNames:  allColorNames,
	}
}
