package core

import (
	"testing"
	"time"

	"github.com/yrichika/gest/pkg/gt"
	"github.com/yrichika/gfaker/pkg/fk/common/util"
	"github.com/yrichika/gfaker/pkg/fk/testutil"
)

func TestRandTime(testingT *testing.T) {
	randTime := NewRandTime(util.RandSeed())

	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)

	tPF := gt.CreateTest(testingT)
	tPF.Describe("PastFuture", func() {
		tPF.It("should return any random past/future time", func() {
			r := randTime.PastFuture()

			gt.Expect(tPF, &r).ToBe_(gt.TimeBetween(past30Years), future30Years)
		})
	})

	tPFr := gt.CreateTest(testingT)
	tPFr.Describe("PastFrom", func() {
		tPFr.It("should return any random past time from the given time", func() {
			r := randTime.PastFrom(past30Years)

			gt.Expect(tPFr, &r).ToBe_(gt.TimeBetween(past30Years), time.Now())
		})
	})

	tP := gt.CreateTest(testingT)
	tP.Describe("Past", func() {
		tP.It("should return any random past time", func() {
			r := randTime.Past()

			gt.Expect(tP, &r).ToBe_(gt.After, past30Years)
		})
	})

	tFT := gt.CreateTest(testingT)
	tFT.Describe("FutureTo", func() {
		tFT.It("should return any random future time to the given time", func() {
			r := randTime.FutureTo(future30Years)

			gt.Expect(tFT, &r).ToBe_(gt.TimeBetween(time.Now()), future30Years)
		})
	})

	tF := gt.CreateTest(testingT)
	tF.Describe("Future", func() {
		tF.It("should return any random future time", func() {
			r := randTime.Future()

			gt.Expect(tF, &r).ToBe_(gt.Before, future30Years)
		})
	})

	tTR := gt.CreateTest(testingT)
	tTR.Describe("TimeRange", func() {
		tTR.It("should return any random time between the given range", func() {
			r := randTime.TimeRange(past30Years, future30Years)

			gt.Expect(tTR, &r).ToBe_(gt.TimeBetween(past30Years), future30Years)
		})
	})

	tD := gt.CreateTest(testingT)
	tD.Describe("Duration", func() {
		r := randTime.Duration()
		// check output
		testutil.Output("RandTime.Duration", r)
	})

	tDMSec := gt.CreateTest(testingT)
	tDMSec.Describe("DurationMilliSec", func() {
		r := randTime.DurationMilliSec()
		// check output
		testutil.Output("RandTime.DurationMilliSec", r)
	})

	tDSec := gt.CreateTest(testingT)
	tDSec.Describe("DurationSec", func() {
		r := randTime.DurationSec()
		// check output
		testutil.Output("RandTime.DurationSec", r)
	})

	tDMin := gt.CreateTest(testingT)
	tDMin.Describe("DurationMin", func() {
		r := randTime.DurationMin()
		// check output
		testutil.Output("RandTime.DurationMin", r)
	})

	tDHour := gt.CreateTest(testingT)
	tDHour.Describe("DurationHour", func() {
		r := randTime.DurationHour()
		// check output
		testutil.Output("RandTime.DurationHour", r)
	})

	tDTo := gt.CreateTest(testingT)
	tDTo.Describe("DurationTo", func() {
		r := randTime.DurationTo(1 * time.Second)
		gt.Expect(tDTo, &r).ToBe_(gt.LessThanOrEq, 1*time.Second)
	})

	tDRange := gt.CreateTest(testingT)
	tDRange.Describe("DurationRange", func() {
		r := randTime.DurationRange(1*time.Second, 2*time.Second)
		gt.Expect(tDRange, &r).ToBe_(gt.Between(1*time.Second), 2*time.Second)
	})
}
