package engine

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// CronTrigger matches the classic 5-field cron format:
// minute hour day-of-month month day-of-week
// Only "*" and exact numbers are supported — no ranges, lists or
// steps yet. Enough to prove the matching logic works.
type CronTrigger struct {
	Expr string

	minute, hour, day, month, weekday cronField
	ready                             bool
}

type cronField struct {
	wildcard bool
	value    int
}

func parseCronField(s string) (cronField, error) {
	if s == "*" {
		return cronField{wildcard: true}, nil
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return cronField{}, fmt.Errorf("cron: bad field %q: %w", s, err)
	}
	return cronField{value: n}, nil
}

func (f cronField) matches(v int) bool {
	return f.wildcard || f.value == v
}

func (t *CronTrigger) init() error {
	if t.ready {
		return nil
	}
	parts := strings.Fields(t.Expr)
	if len(parts) != 5 {
		return fmt.Errorf("cron: expected 5 fields, got %d in %q", len(parts), t.Expr)
	}
	targets := []*cronField{&t.minute, &t.hour, &t.day, &t.month, &t.weekday}
	for i, p := range parts {
		f, err := parseCronField(p)
		if err != nil {
			return err
		}
		*targets[i] = f
	}
	t.ready = true
	return nil
}

// Watch polls once a second until the current time matches the
// expression. Fine at this scale — a real scheduler would compute
// the next match time directly instead of polling.
func (t *CronTrigger) Watch() error {
	if err := t.init(); err != nil {
		return err
	}
	for {
		now := time.Now()
		if t.minute.matches(now.Minute()) &&
			t.hour.matches(now.Hour()) &&
			t.day.matches(now.Day()) &&
			t.month.matches(int(now.Month())) &&
			t.weekday.matches(int(now.Weekday())) {
			return nil
		}
		time.Sleep(time.Second)
	}
}
