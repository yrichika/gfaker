package lorem

import (
	"strings"

	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
)

type Lorem struct {
	rand *core.Rand
	data *provider.Lorems
}

func New(rand *core.Rand, global *provider.Global) *Lorem {
	return &Lorem{
		rand,
		global.Lorems,
	}
}

func (l *Lorem) Word() string {
	return l.rand.Slice.StrElem(l.data.Words)
}

// 複数の文字を配列として返す
// num: 1以上の整数
func (l *Lorem) WordSliceFixedLength(num int) []string {
	if num < 1 {
		num = 1
	}
	var words []string
	for i := 0; i < num; i++ {
		words = append(words, l.Word())
	}
	return words
}

// maxNum: 2以上の整数
func (l *Lorem) WordSlice(maxNum int) []string {
	if maxNum < 2 {
		maxNum = 2
	}
	wordNum := l.rand.Num.IntBt(1, maxNum)
	return l.WordSliceFixedLength(wordNum)
}

// 複数の文字を1つの文字列として返す
func (l *Lorem) Words(num int) string {
	words := l.WordSlice(num)
	return strings.Join(words, " ")
}

// 指定した文字数でランダムな文字列を返す
func (l *Lorem) SentenceFixedLength(wordNum int) string {
	words := l.WordSliceFixedLength(wordNum)
	words[0] = util.CapFirstLetter(words[0])
	return strings.Join(words, " ") + "."
}

// WORKING: 以下の関数からやりなおし

// 文字数の最大値を指定して、それ以下の文字数でランダムな文字列を返す
func (l *Lorem) Sentence(maxWordCount int) string {
	if maxWordCount < 2 {
		maxWordCount = 2
	}
	wordNum := l.rand.Num.IntBt(1, maxWordCount)
	return l.SentenceFixedLength(wordNum)
}

func (l *Lorem) SentenceSliceFixedLength(sentenceNum int, wordMaxNum int) []string {
	var sentences []string
	for i := 0; i < sentenceNum; i++ {
		wordNum := l.rand.Num.IntBt(1, wordMaxNum)
		sentences = append(sentences, l.SentenceFixedLength(wordNum))
	}
	return sentences
}

// 複数の文を配列で返す
func (l *Lorem) SentenceSlice(sentenceMaxNum int, wordMaxNum int) []string {
	if sentenceMaxNum < 2 {
		sentenceMaxNum = 2
	}
	if wordMaxNum < 2 {
		wordMaxNum = 2
	}
	sentenceNum := l.rand.Num.IntBt(1, sentenceMaxNum)
	return l.SentenceSliceFixedLength(sentenceNum, wordMaxNum)
}

// 複数の文を1つの文字列で返す
func (l *Lorem) Sentences(sentenceMaxNum int, wordMaxNum int) string {

	sentences := l.SentenceSlice(sentenceMaxNum, wordMaxNum)
	return strings.Join(sentences, " ")
}

// TEST:
// paragraphNum must be greater than or eq 1
// sentenceMaxNum must be greater than or eq 2
func (l *Lorem) ParagraphSliceFixedLength(paragraphNum int, sentenceMaxNum int) []string {
	if paragraphNum < 1 {
		paragraphNum = 1
	}
	if sentenceMaxNum < 2 {
		sentenceMaxNum = 2

	}
	var paragraphs []string
	for i := 0; i < paragraphNum; i++ {
		wordNum := l.rand.Num.IntBt(1, 30)
		sentenceNum := l.rand.Num.IntBt(1, sentenceMaxNum)
		paragraphs = append(paragraphs, l.Sentences(sentenceNum, wordNum))
	}
	return paragraphs
}

// TEST:
// func (l *Lorem) ParagraphSlice(paragraphMaxNum int, sentenceMaxNum int) []string {
// 	if paragraphMaxNum < 2 {
// 		paragraphMaxNum = 2
// 	}
// 	if sentenceMaxNum < 2 {
// 		sentenceMaxNum = 2
// 	}
// 	paragraphNum := l.rand.Num.IntBt(1, paragraphMaxNum)
// 	return l.ParagraphSliceFixedLength(paragraphNum, sentenceMaxNum)
// }

// TEST:
// func (l *Lorem) Paragraphs(paragraphMaxNum int, sentenceMaxNum int) string {
// 	paragraphs := l.ParagraphSlice(paragraphMaxNum, sentenceMaxNum)
// 	return strings.Join(paragraphs, "\n") // TODO: 改行は1つでいいか、2つでいいか試す
// }

// text
