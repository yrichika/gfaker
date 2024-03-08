package ja_JP

import "github.com/yrichika/gfaker/pkg/fk/provider"

func CreateAddresses() *provider.Addresses {
	return &provider.Addresses{
		CitySuffixes:            CitySuffixes,
		WardSuffixes:            WardSuffixes,
		WardNames:               WardNames,
		CityNames:               CityNames,
		AreaNames:               AreaNames,
		AreaNumbers:             AreaNumbers,
		BuildingNumbers:         BuildingNumbers,
		Postcodes:               Postcodes,
		Prefectures:             Prefectures,
		Countries:               Countries,
		CityFormats:             CityFormats,
		SecondaryAddressFormats: SecondaryAddressFormats,
		AddressFormats:          AddressFormats,
	}
}

var CitySuffixes = []string{"市", "町", "村"}

var WardSuffixes = []string{"区"}

var CityNames = []string{
	"札幌", "函館", "室蘭", "浦河", "旭川", "留萌", "稚内", "網走", "帯広",
	"釧路", "根室", "江別", "千歳", "札幌", "岩見沢", "滝川", "深川", "小樽",
	"倶知安", "函館", "長万部", "江差", "苫小牧", "室蘭", "浦河", "旭川", "士別",
	"名寄", "留萌", "稚内", "網走", "北見", "紋別", "帯広", "釧路", "根室", "日高", "富良野",
	"枝幸", "弟子屈", "恵庭北広島", "石狩", "当別", "月形", "定山渓", "中山峠", "浜益",
	"芦別", "夕張", "三川", "長沼", "美唄", "砂川", "赤平", "余市", "喜茂別", "岩内",
	"寿都", "八雲", "松前", "森", "熊石", "せたな", "登別", "洞爺湖町", "白老", "伊達",
	"洞爺湖", "新ひだか", "えりも", "上川", "音威子府", "美深", "羽幌", "浜頓別", "豊富", "留辺蘂",
	"美幌", "湧別", "遠軽", "興部", "小清水", "斜里", "端野", "清水", "足寄", "本別",
	"広尾", "浦幌", "幕別", "上士幌", "陸別", "白糠", "標茶", "厚岸", "別海", "中標津",
	"標津", "羅臼", "大曲", "栗山", "碧水", "夕張紅葉山", "銭函", "沼ノ端", "常呂", "川湯",
	"日高門別", "北竜", "恵山", "北斗", "安平", "大空", "阿寒湖温泉", "青森", "青森", "八戸",
	"弘前", "十和田", "むつ", "五所川原", "三沢", "黒石", "藤崎", "平川", "野辺地", "十和田湖",
	"七戸", "南部", "三戸", "大間", "龍飛", "鰺ヶ沢", "外ヶ浜", "つがる", "五戸", "おいらせ",
	"盛岡", "盛岡", "一関", "釜石", "北上", "宮古", "久慈", "二戸", "大船渡", "花巻",
	"奥州", "八幡平市", "軽米", "岩泉", "田野畑", "遠野", "平泉", "陸前高田", "岩手町", "滝沢",
	"仙台", "石巻", "大崎", "気仙沼", "白石", "仙台駅", "県庁市役所", "泉中央", "角田", "岩沼",
	"栗原", "松島", "鳴子温泉", "涌谷", "登米", "作並", "大河原", "加美", "女川", "南三陸",
	"長町", "苦竹", "秋田", "秋田", "大館", "能代", "由利本荘", "大仙", "横手", "湯沢",
	"鹿角", "北秋田", "角館", "協和", "田沢湖", "男鹿", "山形", "米沢", "新庄",
	"酒田", "鶴岡", "南陽", "長井", "庄内", "金山", "天童", "寒河江", "東根", "尾花沢",
	"福島", "郡山", "白河", "会津若松", "いわき", "南相馬", "相馬", "南会津", "二本松", "伊達",
	"川俣", "石川", "三春", "小野", "田村", "棚倉", "矢祭", "猪苗代", "西会津", "会津坂下",
	"三島", "只見", "双葉", "浪江", "須賀川", "喜多方", "会津美里", "本宮", "国見", "天栄",
	"下郷", "北塩原", "金山", "昭和", "泉崎", "矢吹", "鮫川", "玉川", "平田", "古殿",
}

var WardNames = []string{"中央", "北", "東", "南", "西", "中", "港"}

var AreaNames = []string{"TODO:"}

var AreaNumbers = []string{
	"#丁目#番地#", "#丁目#番地##", "#丁目#番地###",
	"#丁目##番地#", "#丁目##番地##", "#丁目##番地###",
	"#丁目###番地#", "#丁目###番地##", "#丁目###番地###",
	"#-#-#", "#-#-##", "#-#-###",
	"#-##-#", "#-##-##", "#-##-###",
	"#-###-#", "#-###-##", "#-###-###",
}

var Postcodes = []string{"###-####", "#######"}

var BuildingNumbers = []string{"#0#", "###", "#?", "####", "###?"}

var Countries = []string{
	"アフガニスタン", "アルバニア", "アルジェリア", "アメリカ領サモア", "アンドラ", "アンゴラ", "アンギラ", "南極大陸", "アンティグアバーブーダ", "アルゼンチン", "アルメニア", "アルバ", "オーストラリア", "オーストリア", "アゼルバイジャン",
	"バハマ", "バーレーン", "バングラデシュ", "バルバドス", "ベラルーシ", "ベルギー", "ベリーズ", "ベナン", "バミューダ島", "ブータン", "ボリビア", "ボスニア・ヘルツェゴビナ", "ボツワナ", "ブーベ島", "ブラジル", "イギリス領インド洋地域", "イギリス領ヴァージン諸島", "ブルネイ", "ブルガリア", "ブルキナファソ", "ブルンジ",
	"カンボジア", "カメルーン", "カナダ", "カーボベルデ", "ケイマン諸島", "中央アフリカ共和国", "チャド", "チリ", "中国", "クリスマス島", "ココス諸島", "コロンビア", "コモロ", "コンゴ共和国", "クック諸島", "コスタリカ", "コートジボワール", "クロアチア", "キューバ", "キプロス共和国", "チェコ共和国",
	"デンマーク", "ジブチ共和国", "ドミニカ国", "ドミニカ共和国",
	"エクアドル", "エジプト", "エルサルバドル", "赤道ギニア共和国", "エリトリア", "エストニア", "エチオピア",
	"フェロー諸島", "フォークランド諸島", "フィジー共和国", "フィンランド", "フランス", "フランス領ギアナ", "フランス領ポリネシア", "フランス領極南諸島",
	"ガボン", "ガンビア", "ジョージア", "ドイツ", "ガーナ", "ジブラルタル", "ギリシャ", "グリーンランド", "グレナダ", "グアドループ", "グアム", "グアテマラ", "ガーンジー", "ギニア", "ギニアビサウ", "ガイアナ",
	"ハイチ", "ハード島とマクドナルド諸島", "バチカン市国", "ホンジュラス", "香港", "ハンガリー",
	"アイスランド", "インド", "インドネシア", "イラン", "イラク", "アイルランド共和国", "マン島", "イスラエル", "イタリア",
	"ジャマイカ", "日本", "ジャージー島", "ヨルダン",
	"カザフスタン", "ケニア", "キリバス", "朝鮮", "韓国", "クウェート", "キルギス共和国",
	"ラオス人民民主共和国", "ラトビア", "レバノン", "レソト", "リベリア", "リビア国", "リヒテンシュタイン", "リトアニア", "ルクセンブルク",
	"マカオ", "マケドニア共和国", "マダガスカル", "マラウィ", "マレーシア", "モルディブ", "マリ", "マルタ共和国", "マーシャル諸島", "マルティニーク", "モーリタニア・イスラム共和国", "モーリシャス", "マヨット", "メキシコ", "ミクロネシア連邦", "モルドバ共和国", "モナコ公国", "モンゴル", "モンテネグロ共和国", "モントセラト", "モロッコ", "モザンビーク", "ミャンマー",
	"ナミビア", "ナウル", "ネパール", "オランダ領アンティル", "オランダ", "ニューカレドニア", "ニュージーランド", "ニカラグア", "ニジェール", "ナイジェリア", "ニース", "ノーフォーク島", "北マリアナ諸島", "ノルウェー",
	"オマーン",
	"パキスタン", "パラオ", "パレスチナ自治区", "パナマ", "パプアニューギニア", "パラグアイ", "ペルー", "フィリピン", "ピトケアン諸島", "ポーランド", "ポルトガル", "プエルトリコ",
	"カタール",
	"レユニオン", "ルーマニア", "ロシア", "ルワンダ",
	"サン・バルテルミー島", "セントヘレナ", "セントクリストファー・ネイビス連邦", "セントルシア", "セント・マーチン島", "サンピエール島・ミクロン島", "セントビンセント・グレナディーン", "サモア", "サンマリノ", "サントメプリンシペ", "サウジアラビア", "セネガル", "セルビア", "セイシェル", "シエラレオネ", "シンガポール", "スロバキア", "スロベニア", "ソロモン諸島", "ソマリア", "南アフリカ共和国", "サウスジョージア・サウスサンドウィッチ諸島", "スペイン", "スリランカ", "スーダン", "スリナム", "スヴァールバル諸島およびヤンマイエン島", "スワジランド王国", "スウェーデン", "スイス", "シリア",
	"台湾", "タジキスタン共和国", "タンザニア", "タイ", "東ティモール", "トーゴ", "トケラウ", "トンガ", "トリニダード・トバゴ", "チュニジア", "トルコ", "トルクメニスタン", "タークス・カイコス諸島", "ツバル",
	"ウガンダ", "ウクライナ", "アラブ首長国連邦", "イギリス", "アメリカ合衆国", "合衆国領有小離島", "アメリカ領ヴァージン諸島", "ウルグアイ", "ウズベキスタン",
	"バヌアツ", "ベネズエラ", "ベトナム",
	"ウォリス・フツナ", "西サハラ",
	"イエメン",
	"ザンビア", "ジンバブエ",
}
var Prefectures = []string{
	"北海道",
	"青森県", "岩手県", "宮城県", "秋田県", "山形県", "福島県",
	"茨城県", "栃木県", "群馬県", "埼玉県", "千葉県", "東京都", "神奈川県",
	"新潟県", "富山県", "石川県", "福井県", "山梨県", "長野県", "岐阜県", "静岡県", "愛知県",
	"三重県", "滋賀県", "京都府", "大阪府", "兵庫県", "奈良県", "和歌山県",
	"鳥取県", "島根県", "岡山県", "広島県", "山口県",
	"徳島県", "香川県", "愛媛県", "高知県",
	"福岡県", "佐賀県", "長崎県", "熊本県", "大分県", "宮崎県", "鹿児島県",
	"沖縄県",
}

var CityFormats = []string{"{{.CityName}}{{.CitySuffix}}"}
var WardFormats = []string{"{{.WardName}}{{.WardSuffix}}"}
var AreaFormats = []string{"{{.AreaName}}{{.AreaNumber}}"}
var SecondaryAddressFormats = []string{
	"ハイツ{{.LastName}}{{.BuildingNumber}}号",
	"コーポ{{.LastName}}{{.BuildingNumber}}号",
}
var AddressFormats = []string{
	"{{.Postcode}}  {{.Prefecture}}{{.City}}{{.Area}}",
	"{{.Postcode}}  {{.Prefecture}}{{.City}}{{.Ward}}{{.Area}}",
	"{{.Postcode}}  {{.Prefecture}}{{.City}}{{.Ward}}{{.Area}} {{.SecondaryAddress}}",
}