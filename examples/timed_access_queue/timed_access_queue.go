package timedAccessQueue

import (
	"time"

	"github.com/DangeL187/timed_access"
)

type Queue[T any] interface {
	Push(item T)
	Pop() (T, bool)
	Len() int
}

type TimedAccessQueue[T any, Q Queue[T]] struct {
	queue Q

	timedAccess.TimedAccess[T]
}

func (t *TimedAccessQueue[T, Q]) Push(item T) {
	timedAccess.DoInSafeIntervalVoid(t.IsInSafeInterval, func() {
		t.queue.Push(item)
	})
}

func (t *TimedAccessQueue[T, Q]) Pop() (T, bool) {
	return timedAccess.DoInSafeInterval2(t.IsInSafeInterval, func() (T, bool) {
		return t.queue.Pop()
	})
}

func (t *TimedAccessQueue[T, Q]) Len() int {
	return t.queue.Len()
}

func NewTimedAccessQueue[T any, Q Queue[T]](queue Q, spinPeriod, intervalSize time.Duration) *TimedAccessQueue[T, Q] {
	t := &TimedAccessQueue[T, Q]{
		queue: queue,
	}

	t.SetPeriod(spinPeriod)
	t.SetIntervalSize(intervalSize)

	return t
}
