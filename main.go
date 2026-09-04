package main

import (
	"time"

	"github.com/cheemney/triggr/engine"
)

func main() {
	rule := &engine.Rule{
		Name:    "heartbeat",
		Trigger: &engine.IntervalTrigger{Interval: 2 * time.Second},
		Action:  &engine.LogAction{Message: "heartbeat: rule fired"},
	}

	if err := rule.Run(); err != nil {
		panic(err)
	}
}
