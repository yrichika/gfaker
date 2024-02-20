package core

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yrichika/gfaker/pkg/fk/common/log"
)

// time パッケージをランダムに扱うもの
// datetimeとの違いは、datetimeはtimeを使って、様々な形式で日時を返すが、timeはtime.Timeを返す

type RandTime struct {
	rand *rand.Rand
}

func NewRandTime(rand *rand.Rand) *RandTime {
	return &RandTime{
		rand,
	}
}

// 30年前から30年後までのランダムな日時を返す
func (r *RandTime) PastFuture() time.Time {
	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)
	return r.TimeRange(past30Years, future30Years)
}

// 指定した過去から今までのランダムな日時を返す
func (r *RandTime) PastFrom(from time.Time) time.Time {
	if from.After(time.Now()) {
		errMsg := fmt.Sprintf("Invalid past time: from=%#v", from)
		log.WrongUsage(errMsg, 1)
		return time.Time{}
	}
	return r.TimeRange(from, time.Now())
}

// 30年前から今までのランダムな日時を返す
func (r *RandTime) Past() time.Time {
	past30Years := time.Now().Add(-30 * 365 * 24 * time.Hour)
	return r.PastFrom(past30Years)
}

// 現在から指定した日時までのランダムな日時を返す
func (r *RandTime) FutureTo(to time.Time) time.Time {
	if to.Before(time.Now()) {
		errMsg := fmt.Sprintf("Invalid future time: to=%#v", to)
		log.WrongUsage(errMsg, 1)
		return time.Time{}
	}
	return r.TimeRange(time.Now(), to)
}

// 今から30年後までのランダムな日時を返す
func (r *RandTime) Future() time.Time {
	future30Years := time.Now().Add(30 * 365 * 24 * time.Hour)
	return r.FutureTo(future30Years)
}

// 指定された範囲の中でランダムな日時を返す
func (r *RandTime) TimeRange(from time.Time, to time.Time) time.Time {
	if from.After(to) {
		errMsg := fmt.Sprintf("Invalid range: from=%#v, to=%#v", from, to)
		log.WrongUsage(errMsg, 1)
		return time.Time{}
	}
	diff := to.Sub(from)
	randomDiff := time.Duration(r.rand.Int63n(int64(diff)))
	return from.Add(randomDiff)
}

func (r *RandTime) Duration() time.Duration {
	return time.Duration(r.rand.Int63())
}

// 1秒以下のランダムなミリ秒を返す
func (r *RandTime) DurationMilliSec() time.Duration {
	return r.DurationTo(1000 * time.Millisecond)
}

// 1分以下のランダムな秒を返す
func (r *RandTime) DurationSec() time.Duration {
	return r.DurationTo(60 * time.Second)
}

// 1時間以下のランダムな分を返す
func (r *RandTime) DurationMin() time.Duration {
	return r.DurationTo(60 * time.Minute)
}

// 1日以下のランダムな時間を返す
func (r *RandTime) DurationHour() time.Duration {
	return r.DurationTo(24 * time.Hour)
}

// 指定した時間以下のランダムな時間を返す
func (r *RandTime) DurationTo(to time.Duration) time.Duration {
	return time.Duration(r.rand.Int63n(int64(to)))
}

// 指定した範囲の中でランダムな時間を返す
func (r *RandTime) DurationRange(from time.Duration, to time.Duration) time.Duration {
	if from > to {
		errMsg := fmt.Sprintf("Invalid range: from=%#v, to=%#v", from, to)
		log.WrongUsage(errMsg, 1)
		return 0
	}
	diff := to - from
	randomDiff := time.Duration(r.rand.Int63n(int64(diff)))
	return from + randomDiff
}
