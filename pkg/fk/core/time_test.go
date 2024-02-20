package core

import (
	"math/rand"
	"testing"
	"time"

	"github.com/yrichika/gest/pkg/gt"
)

func TestRandTime(testingT *testing.T) {
	t := gt.CreateTest(testingT)
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))
	randTime := NewRandTime(rand)
	t.Describe("Past", func() {
		t.It("should return any random past time", func() {
			r := randTime.Past()

			expectedMaxPast := time.Now().Add(-30 * 365 * 24 * time.Hour)
			gt.Expect(t, &r).ToBe_(gt.After, expectedMaxPast)
		})
	})
}
