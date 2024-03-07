package fk

import (
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/generator/address"
	"github.com/yrichika/gfaker/pkg/fk/generator/color"
	"github.com/yrichika/gfaker/pkg/fk/generator/person"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
)

type Faker struct {
	Rand    *core.Rand
	Person  *person.Person
	Color   *color.Color
	Address *address.Address
	// TODO: Faker/Factoryの $defaultProvidersの変数にあるものをここに入れる
	// Company *provider.Company
	// ...et

}

// REF: https://fakerphp.github.io/

func Create() *Faker {
	localized := en_US.New()
	return CreateWithLocale(localized)
}

func CreateWithLocale(localized *provider.Localized) *Faker {
	coreRand := core.NewRand(util.RandSeed())
	global := global.New()
	return &Faker{
		Rand:   coreRand,
		Person: person.New(coreRand, localized),
		Color:  color.New(coreRand, global),
	}
}
