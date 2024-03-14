package fk

import (
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/generator/address"
	"github.com/yrichika/gfaker/pkg/fk/generator/barcode"
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
	Barcode *barcode.Barcode
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
		Rand:    coreRand,
		Barcode: barcode.New(coreRand),
		Color:   color.New(coreRand, global),
		Person:  person.New(coreRand, localized),
		Address: address.New(coreRand, localized),
	}
}
