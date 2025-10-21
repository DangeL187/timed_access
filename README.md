# About
**TimedAccess** is a Go library for time-controlled execution of functions. It allows running actions only within predefined **safe intervals** which helps to:
- Smooth out resource usage
- Synchronize operations with external timers
- Avoid excessive calls in high-throughput systems

> [!Note]
> **TimedAccess** does not provide thread safety by itself. For concurrent use, combine it with CAS, mutex, or thread-safe data structures inside your actions.

While **TimedAccess** does not guarantee full thread safety, in scenarios with low-frequency writes and occasional random accesses, **TimedAccess** can provide a practically thread-safe experience.

For example, when one thread modifies the data once every **N** seconds and another thread accesses it at **random** times, wrapping operations with **TimedAccess** ensures that the second thread only executes actions during safe intervals.

This approach minimizes the risk of collisions and can reduce the need for heavy locks or CAS in such situations.

# How it Works
**TimedAccess** stores:
- **period** - the interval in nanoseconds between successive accesses by a primary thread (e.g., Thread 1)
- **intervalSize** - the maximum duration in nanoseconds that the primary thread may hold or modify the data during a single access
- **startTime** - the reference start time from which intervals are calculated

**IsInSafeInterval** returns (ok, nextSleep):
- **ok** = true → current time is inside a safe interval
- **nextSleep** → duration to wait before retrying if currently outside the safe interval

Actions executed via **DoInSafeInterval** are guaranteed to run only during these safe periods, minimizing collisions with the primary thread

---

Before using TimedAccess, you should set the start time using the SetStartTime function. The start time represents the first moment when Thread 1 will access the data.

For example, if you have a ticker like this:
```go
ticker := time.NewTicker(time.Second)
```
then the start time should be set to:
```go
ta.SetStartTime(time.Now().Add(time.Second))
```
This ensures that the safe and forbidden intervals in TimedAccess **align correctly with the first tick of your thread**.

# Example

Legend:
- `[====]` → Safe interval (operations allowed)
- `[----]` → Unsafe interval (operations blocked)

Parameters:
- period = 1s
- intervalSize = 0.1s

Timeline (first tick of Thread 1 at 1s):
```
0s        0.1s      0.5s      1s        1.1s      1.5s      2s        2.1s      2.5s      3s
|         |         |         |         |         |         |         |         |         |
|=============================----------====================----------====================-
```
