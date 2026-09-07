package engine

import (
	"context"
	"log"
	"sync"
)

// Engine runs a set of rules concurrently. Each rule loops forever —
// watch, execute, watch again — until ctx is cancelled.
type Engine struct {
	rules []*Rule
}

func NewEngine(rules []*Rule) *Engine {
	return &Engine{rules: rules}
}

// Run blocks until every rule's loop has stopped.
func (e *Engine) Run(ctx context.Context) {
	var wg sync.WaitGroup
	for _, r := range e.rules {
		wg.Add(1)
		go func(r *Rule) {
			defer wg.Done()
			e.loop(ctx, r)
		}(r)
	}
	wg.Wait()
}

func (e *Engine) loop(ctx context.Context, r *Rule) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		// note: ctx is only checked between iterations, not while
		// Watch() itself is blocked (e.g. mid-sleep) — fine for now,
		// worst case shutdown is delayed by one trigger cycle.
		if err := r.Run(); err != nil {
			log.Printf("rule %q failed: %v", r.Name, err)
			return
		}
	}
}
