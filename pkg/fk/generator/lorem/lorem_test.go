package lorem

import (
	"testing"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/core"
	"github.com/yrichika/gfaker/pkg/fk/provider"
	"github.com/yrichika/gfaker/pkg/fk/provider/global"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestLorem(testingT *testing.T) {
	coreRand := core.NewRand(util.RandSeed())
	global := &provider.Global{
		Lorems: global.CreateLorems(),
	}

	lorem := New(coreRand, global)

	t := gt.CreateTest(testingT)
	t.Describe("Lorem", func() {
		t.Test("Word should return a word", func() {
			r := lorem.Word()
			gt.Expect(t, &r).ToBeIn(lorem.data.Words)
		})

		t.Test("WordSliceFixedLength should return a slice of words with fixed length", func() {
			r := lorem.WordSliceFixedLength(5)
			for _, v := range r {
				gt.Expect(t, &v).ToBeIn(lorem.data.Words)
			}
		})

		t.Test("WordSlice should return a slice of words", func() {
			r := lorem.WordSlice(5)
			for _, v := range r {
				gt.Expect(t, &v).ToBeIn(lorem.data.Words)
			}
		})

		t.Test("Words should return a string of words", func() {
			r := lorem.Words(5)
			testutil.Output("Lorem.Words", r)
		})

		t.Test("SentenceFixedLength should return a sentence with fixed length", func() {
			r := lorem.SentenceFixedLength(6)
			testutil.Output("Lorem.SentenceFixedLength", r)
		})

		// t.Test("Sentence should return a sentence", func() {
		// 	r := lorem.Sentence(6)
		// 	testutil.Output("Lorem.Sentence", r)
		// })

		// t.Test("SentencesFixedLength should return sentences with fixed length", func() {
		// 	r := lorem.SentenceSliceFixedLength(3, 6)
		// 	testutil.Output("Lorem.SentencesFixedLength", r)
		// })

		// t.Test("SentenceSlice should return slices of sentences", func() {
		// 	r := lorem.SentenceSlice(3, 6)
		// 	testutil.Output("Lorem.SentenceSlice", r)
		// })

		// t.Test("Sentences should return a string of sentences", func() {
		// 	r := lorem.Sentences(3, 6)
		// 	testutil.Output("Lorem.Sentences", r)
		// })

		// t.Test("ParagraphFixedLength should return a paragraph with fixed length of sentences", func() {
		// 	r := lorem.ParagraphSliceFixedLength(6, 6)
		// 	testutil.Output("Lorem.ParagraphFixedLength", r)
		// })
	})

}
