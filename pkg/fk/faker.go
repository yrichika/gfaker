package fk

import (
	"math/rand"
	"time"

	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/base/person"
	"github.com/yrichika/gfaker/pkg/fk/provider/locale/en_US"
)

type Faker struct {
	Rand   *core.Rand
	Person *person.Person
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
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	coreRand := core.NewRand(rand)
	return &Faker{
		Rand:   coreRand,
		Person: person.New(coreRand, localized),
	}
}
