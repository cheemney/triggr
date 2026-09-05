package main

import (
	"fmt"
	"time"

	"github.com/cheemney/triggr/engine"
)

func main() {
	heartbeat := &engine.Rule{
		Name:    "heartbeat",
		Trigger: &engine.IntervalTrigger{Interval: 2 * time.Second},
		Action:  &engine.LogAction{Message: "heartbeat: rule fired"},
	}
	if err := heartbeat.Run(); err != nil {
		panic(err)
	}

	// wildcard everything except the current hour, so this fires
	// almost immediately but still proves field matching works —
	// change the hour and it won't fire.
	hour := time.Now().Hour()
	cronDemo := &engine.Rule{
		Name:    "cron-demo",
		Trigger: &engine.CronTrigger{Expr: fmt.Sprintf("* %d * * *", hour)},
		Action:  &engine.LogAction{Message: "cron: matched current hour"},
	}
	if err := cronDemo.Run(); err != nil {
		panic(err)
	}
}
