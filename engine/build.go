package engine

import (
	"fmt"
	"time"

	"github.com/cheemney/triggr/config"
)

// BuildRule turns a parsed config.RuleDecl into a runnable Rule.
// Only "interval" triggers and "shell"/"log" actions are wired up so
// far — cron, webhook and http will get plugged in as they're needed.
func BuildRule(decl config.RuleDecl) (*Rule, error) {
	trigger, err := buildTrigger(decl.Trigger)
	if err != nil {
		return nil, err
	}
	action, err := buildAction(decl.Action)
	if err != nil {
		return nil, err
	}
	return &Rule{Name: decl.Name, Trigger: trigger, Action: action}, nil
}

func buildTrigger(spec config.Spec) (Trigger, error) {
	switch spec.Kind {
	case "interval":
		if len(spec.Args) != 1 {
			return nil, fmt.Errorf("interval trigger: expected 1 arg, got %d", len(spec.Args))
		}
		d, err := time.ParseDuration(spec.Args[0])
		if err != nil {
			return nil, fmt.Errorf("interval trigger: %w", err)
		}
		return &IntervalTrigger{Interval: d}, nil
	case "github_release":
		if len(spec.Args) != 3 {
			return nil, fmt.Errorf("github_release trigger: expected 3 args (owner, repo, interval), got %d", len(spec.Args))
		}
		d, err := time.ParseDuration(spec.Args[2])
		if err != nil {
			return nil, fmt.Errorf("github_release trigger: %w", err)
		}
		return &GitHubReleaseTrigger{Owner: spec.Args[0], Repo: spec.Args[1], Interval: d}, nil
	default:
		return nil, fmt.Errorf("unknown trigger kind %q", spec.Kind)
	}
}

func buildAction(spec config.Spec) (Action, error) {
	switch spec.Kind {
	case "shell":
		if len(spec.Args) != 1 {
			return nil, fmt.Errorf("shell action: expected 1 arg, got %d", len(spec.Args))
		}
		return &ShellAction{Command: spec.Args[0]}, nil
	case "log":
		if len(spec.Args) != 1 {
			return nil, fmt.Errorf("log action: expected 1 arg, got %d", len(spec.Args))
		}
		return &LogAction{Message: spec.Args[0]}, nil
	case "discord_webhook":
		if len(spec.Args) != 2 {
			return nil, fmt.Errorf("discord_webhook action: expected 2 args (url, message), got %d", len(spec.Args))
		}
		return &DiscordWebhookAction{WebhookURL: spec.Args[0], Message: spec.Args[1]}, nil
	default:
		return nil, fmt.Errorf("unknown action kind %q", spec.Kind)
	}
}
