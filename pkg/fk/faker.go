package fk

import (
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/base/color"
	"github.com/yrichika/gfaker/pkg/fk/provider/base/person"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/global"
)

type Faker struct {
	Rand   *core.Rand
	Person *person.Person
	Color  *color.Color
	// TODO: Faker/Factoryの $defaultProvidersの変数にあるものをここに入れる
	// Address *provider.Address
	// Color  *provider.Color
	// Company *provider.Company
	// ...et

}

// REF: https://fakerphp.github.io/

func Create() *Faker {
	localized := en_US.New()
	return CreateWithLocale(localized)
}

func CreateWithLocale(localized *base.Localized) *Faker {
	coreRand := core.NewRand(util.RandSeed())
	global := global.New()
	return &Faker{
		Rand:   coreRand,
		Person: person.New(coreRand, localized),
		Color:  color.New(coreRand, global),
	}
}
