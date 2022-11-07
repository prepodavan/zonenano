package timeutil

import "time"

type NowFunc func() time.Time

var Default = time.Now

func Constant(instant time.Time) NowFunc {
	return func() time.Time {
		return instant
	}
}
