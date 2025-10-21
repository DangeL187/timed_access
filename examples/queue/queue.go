package queue

import (
	"sync/atomic"
)

// Queue with data race detection:

type Float32Queue struct {
	data []float32
	lock int32 // 0 = free, 1 = busy
}

func (q *Float32Queue) Push(item float32) {
	if !atomic.CompareAndSwapInt32(&q.lock, 0, 1) {
		panic("DATA RACE DETECTED on Push!")
	}
	defer atomic.StoreInt32(&q.lock, 0)

	q.data = append(q.data, item)
}

func (q *Float32Queue) Pop() (float32, bool) {
	var zero float32

	if !atomic.CompareAndSwapInt32(&q.lock, 0, 1) {
		panic("DATA RACE DETECTED on Pop!")
	}
	defer atomic.StoreInt32(&q.lock, 0)

	if len(q.data) == 0 {
		return zero, false
	}
	item := q.data[0]
	q.data = q.data[1:]
	return item, true
}

func (q *Float32Queue) Len() int {
	if !atomic.CompareAndSwapInt32(&q.lock, 0, 1) {
		panic("DATA RACE DETECTED on Len!")
	}
	defer atomic.StoreInt32(&q.lock, 0)

	return len(q.data)
}
