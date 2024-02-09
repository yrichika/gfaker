package person

import (
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
)

type Person struct {
	rand  *core.Rand
	local *base.Localized
}

func New(rand *core.Rand, local *base.Localized) *Person {
	return &Person{
		rand,
		local,
	}
}

// TODO: randとlocalを使って、ランダムにlocaleで渡したデータを返すようにする

func (p *Person) FirstNameMale() string {
	return p.local.FirstNameMale[p.rand.Num.Int(0, len(p.local.FirstNameMale))]
}

// TODO: ここに、全てのメソッドを定義する。
// en_USでは、カナの名前はないが、そういった場合は、値がnilかチェックして
// log.Printf() でエラーを出すようにする
// 返す値は、"" か、0 など、型に従って、ゼロ値を返すようにする
