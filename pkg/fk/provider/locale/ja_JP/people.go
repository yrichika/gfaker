package ja_JP

import (
	"github.com/yrichika/gfaker/pkg/fk/provider/base"
	"github.com/yrichika/gfaker/pkg/fk/provider/base/person"
)

func CreatePeople() *base.People {
	return &base.People{
		FirstNameMales:       firstNameMales,
		FirstNameFemales:     firstNameFemales,
		LastNames:            lastNames,
		FirstKanaNameMales:   firstKanaNameMales,
		FirstKanaNameFemales: firstKanaNameFemales,
		LastKanaNames:        lastKanaNames,
		MaleNameFormats:      maleNameFormats,
		FemaleNameFormats:    femaleNameFormats,
		CreateNameMale:       createJaJpNameMale,
		CreateNameFemale:     createJaJpNameFemale,
	}
}

var firstNameMales = []string{
	"晃", "篤司", "治", "和也", "京助", "健一", "修平", "翔太", "淳", "聡太郎", "太一", "太郎", "拓真", "翼", "智也",
	"直樹", "直人", "英樹", "浩", "学", "充", "稔", "裕樹", "裕太", "康弘", "陽一", "洋介", "亮介", "涼平", "零",
}

var firstNameFemales = []string{
	"明美", "あすか", "香織", "加奈", "くみ子", "さゆり", "知実", "千代",
	"直子", "七夏", "花子", "春香", "真綾", "舞", "美加子", "幹", "桃子", "結衣", "裕美子", "陽子", "里佳",
}

var lastNames = []string{
	"青田", "青山", "石田", "井高", "伊藤", "井上", "宇野", "江古田", "大垣",
	"加藤", "加納", "喜嶋", "木村", "桐山", "工藤", "小泉", "小林", "近藤",
	"斉藤", "坂本", "佐々木", "佐藤", "笹田", "鈴木", "杉山",
	"高橋", "田中", "田辺", "津田",
	"中島", "中村", "渚", "中津川", "西之園", "野村",
	"原田", "浜田", "廣川", "藤本",
	"松本", "三宅", "宮沢", "村山",
	"山岸", "山口", "山田", "山本", "吉田", "吉本",
	"若松", "渡辺",
}

var firstKanaNameMales = []string{
	"アキラ", "アツシ", "オサム", "カズヤ", "キョウスケ", "ケンイチ", "シュウヘイ", "ショウタ", "ジュン", "ソウタロウ",
	"タイチ", "タロウ", "タクマ", "ツバサ", "トモヤ", "ナオキ", "ナオト", "ヒデキ", "ヒロシ", "マナブ", "ミツル", "ミノル",
	"ユウキ", "ユウタ", "ヤスヒロ", "ヨウイチ", "ヨウスケ", "リョウスケ", "リョウヘイ", "レイ",
}

var firstKanaNameFemales = []string{
	"アケミ", "アスカ", "カオリ", "カナ", "クミコ", "サユリ", "サトミ", "チヨ",
	"ナオコ", "ナナミ", "ハナコ", "ハルカ", "マアヤ", "マイ", "ミカコ", "ミキ", "モモコ", "ユイ", "ユミコ", "ヨウコ", "リカ",
}

var lastKanaNames = []string{
	"アオタ", "アオヤマ", "イシダ", "イダカ", "イトウ", "ウノ", "エコダ", "オオガキ",
	"カノウ", "カノウ", "キジマ", "キムラ", "キリヤマ", "クドウ", "コイズミ", "コバヤシ", "コンドウ",
	"サイトウ", "サカモト", "ササキ", "サトウ", "ササダ", "スズキ", "スギヤマ",
	"タカハシ", "タナカ", "タナベ", "ツダ",
	"ナカジマ", "ナカムラ", "ナギサ", "ナカツガワ", "ニシノソノ", "ノムラ",
	"ハラダ", "ハマダ", "ヒロカワ", "フジモト",
	"マツモト", "ミヤケ", "ミヤザワ", "ムラヤマ",
	"ヤマギシ", "ヤマグチ", "ヤマダ", "ヤマモト", "ヨシダ", "ヨシモト",
	"ワカマツ", "ワタナベ",
}

var maleNameFormats = []string{
	"{{.LastName}} {{.FirstName}}",
}

var femaleNameFormats = []string{
	"{{.LastName}} {{.FirstName}}",
}

type JaJpPersonName struct {
	FirstName string
	LastName  string
}

func createJaJpNameMale(p interface{}) any {
	a := p.(*person.Person)
	return JaJpPersonName{
		FirstName: a.FirstNameMale(),
		LastName:  a.LastName(),
	}
}

func createJaJpNameFemale(p interface{}) any {
	a := p.(*person.Person)
	return JaJpPersonName{
		FirstName: a.FirstNameFemale(),
		LastName:  a.LastName(),
	}
}
