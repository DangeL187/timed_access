package main

import (
	"math/rand"
	"time"

	"example/queue"
	"example/timed_access_queue"
)

type Queue[T any] interface {
	Push(item T)
	Pop() (T, bool)
	Len() int
}

func spin(queue Queue[float32], startTime time.Time, period time.Duration) {
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			timeSinceStart := float32(time.Since(startTime).Nanoseconds()) / 1000000000
			queue.Push(timeSinceStart)
		}
	}
}

func fire(queue Queue[float32], startTime time.Time) {
	timeSinceStart := float32(time.Since(startTime).Nanoseconds()) / 1000000000

	queue.Push(timeSinceStart)
}

func main() {
	period := time.Second
	intervalSize := time.Millisecond * 100

	taq := timedAccessQueue.NewTimedAccessQueue(&queue.Float32Queue{}, period, intervalSize)

	startTime := time.Now().Add(period)

	taq.SetStartTime(startTime)
	go spin(taq, startTime, period)

	for {
		fire(taq, startTime)
		delay := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(delay)
	}

	//taq.Stop()
}
