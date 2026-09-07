package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cheemney/triggr/config"
	"github.com/cheemney/triggr/engine"
)

func main() {
	data, err := os.ReadFile("examples/github-to-discord/rules.conf")
	if err != nil {
		panic(err)
	}

	program, err := config.Parse(string(data))
	if err != nil {
		panic(err)
	}

	var rules []*engine.Rule
	for _, decl := range program.Rules {
		rule, err := engine.BuildRule(decl)
		if err != nil {
			panic(err)
		}
		rules = append(rules, rule)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	fmt.Printf("running %d rule(s), press Ctrl+C to stop\n", len(rules))
	engine.NewEngine(rules).Run(ctx)
}
