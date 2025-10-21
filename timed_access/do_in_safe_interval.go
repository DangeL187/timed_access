package timedAccess

import (
	"time"
)

// DoInSafeIntervalVoid - no args, returns nothing
func DoInSafeIntervalVoid(isSafe func() (bool, time.Duration), action func()) {
	for {
		ok, nextSleep := isSafe()
		if ok {
			action()
			return
		}

		time.Sleep(nextSleep)
	}
}

// DoInSafeInterval - no args, returns 1 value
func DoInSafeInterval[A any](isSafe func() (bool, time.Duration), action func() A) A {
	for {
		ok, nextSleep := isSafe()
		if ok {
			return action()
		}

		time.Sleep(nextSleep)
	}
}

// DoInSafeInterval2 - no args, returns 2 values
func DoInSafeInterval2[A, B any](isSafe func() (bool, time.Duration), action func() (A, B)) (A, B) {
	for {
		ok, nextSleep := isSafe()
		if ok {
			return action()
		}

		time.Sleep(nextSleep)
	}
}

// DoInSafeIntervalWithArgsVoid - with args, returns nothing
func DoInSafeIntervalWithArgsVoid[Args any](isSafe func() (bool, time.Duration), action func(Args), args Args) {
	for {
		ok, nextSleep := isSafe()
		if ok {
			action(args)
			return
		}

		time.Sleep(nextSleep)
	}
}

// DoInSafeIntervalWithArgs - with args, returns 1 value
func DoInSafeIntervalWithArgs[Args, A any](isSafe func() (bool, time.Duration), action func(Args) A, args Args) A {
	for {
		ok, nextSleep := isSafe()
		if ok {
			return action(args)
		}

		time.Sleep(nextSleep)
	}
}

// DoInSafeIntervalWithArgs2 - with args, returns 2 values
func DoInSafeIntervalWithArgs2[Args, A, B any](isSafe func() (bool, time.Duration), action func(Args) (A, B), args Args) (A, B) {
	for {
		ok, nextSleep := isSafe()
		if ok {
			return action(args)
		}

		time.Sleep(nextSleep)
	}
}
