package usecase

import (
	"context"
	"time"

	"github.com/prepodavan/zonenano/util/timeutil"
)

type UnixTimeCalculator struct {
	now timeutil.NowFunc
}

func NewUnixTimeCalculator() *UnixTimeCalculator {
	return &UnixTimeCalculator{
		now: timeutil.Default,
	}
}

func (calc *UnixTimeCalculator) CalcNanos(_ context.Context, zoneName string, offset int) int64 {
	return calc.now().In(time.FixedZone(zoneName, offset)).UnixNano()
}

func (calc *UnixTimeCalculator) CalcNanosForInstant(
	_ context.Context,
	instant, layout, zoneName string,
	offset int,
) (int64, error) {
	date, err := time.ParseInLocation(layout, instant, time.FixedZone(zoneName, offset))
	if err != nil {
		return 0, err
	}
	return date.UnixNano(), nil
}
