package timedAccess

import (
	"time"
)

type TimedAccess struct {
	intervalSize time.Duration
	period       time.Duration

	startTime time.Time
}

func (t *TimedAccess) SetStartTime(startTime time.Time) {
	t.startTime = startTime
}

func (t *TimedAccess) SetIntervalSize(intervalSize time.Duration) {
	t.intervalSize = intervalSize
}

func (t *TimedAccess) SetPeriod(period time.Duration) {
	t.period = period
}

func (t *TimedAccess) IsInSafeInterval() (bool, time.Duration) {
	if t.period <= 0 || t.startTime.IsZero() {
		return false, 0
	}

	now := time.Now()
	if now.Before(t.startTime) {
		return false, 0
	}

	periodNs := uint64(t.period.Nanoseconds())
	intervalNs := uint64(t.intervalSize.Nanoseconds())
	remainder := uint64(now.Sub(t.startTime).Nanoseconds()) % periodNs

	// If the time elapsed since t.startTime (x) is a multiple of period,
	// or the next multiple of period falls within the range [x, x+t.intervalSize]
	if remainder == 0 || (periodNs-remainder) <= intervalNs {
		return false, time.Duration(periodNs - remainder)
	}

	return true, 0
}
