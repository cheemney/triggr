package engine

import "time"

// IntervalTrigger fires once after a fixed delay. No calendar/cron logic
// yet — that's a later phase, this is just enough to prove the loop works.
type IntervalTrigger struct {
	Interval time.Duration
}

func (t *IntervalTrigger) Watch() error {
	time.Sleep(t.Interval)
	return nil
}
