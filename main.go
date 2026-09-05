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

	fmt.Println("waiting for a POST to http://localhost:8080/fire ...")
	webhookDemo := &engine.Rule{
		Name:    "webhook-demo",
		Trigger: &engine.WebhookTrigger{Addr: ":8080", Path: "/fire"},
		Action:  &engine.LogAction{Message: "webhook: fired"},
	}
	if err := webhookDemo.Run(); err != nil {
		panic(err)
	}

	shellDemo := &engine.Rule{
		Name:    "shell-demo",
		Trigger: &engine.IntervalTrigger{Interval: 1 * time.Second},
		Action:  &engine.ShellAction{Command: "echo hello from shell action"},
	}
	if err := shellDemo.Run(); err != nil {
		panic(err)
	}

	httpDemo := &engine.Rule{
		Name:    "http-demo",
		Trigger: &engine.IntervalTrigger{Interval: 1 * time.Second},
		Action:  &engine.HTTPAction{URL: "https://api.github.com"},
	}
	if err := httpDemo.Run(); err != nil {
		panic(err)
	}
}
