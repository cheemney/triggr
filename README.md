# triggr

*trigger*, minus the vowel — a minimal, dependency-free automation
engine written in Go. If *this* happens, do *that*: same idea as
IFTTT, built from scratch as a learning project.

## How it works

Three types cover everything:

- **Trigger** — blocks until a condition is met (`Watch()`).
- **Action** — runs once a trigger fires (`Execute()`).
- **Rule** — pairs one Trigger with one Action.

`Engine` runs a set of rules concurrently, each looping forever
(watch, execute, repeat) until it's told to stop.

Rules aren't hardcoded — they're written in a small config language
and parsed by a hand-written lexer/parser (`config/`), then turned
into real `Trigger`/`Action` values by `engine/build.go`.

```
rule heartbeat {
    trigger interval 2s
    action log "still alive"
}
```

## What's built in

Triggers: `interval`, `cron` (minute-granularity), `webhook`,
`github_release`.
Actions: `log`, `shell`, `http`, `discord_webhook`.

Adding a new one is just a new Go file plus a case in `build.go` —
the engine and config language don't change.

## State

Triggers that need to remember something between polls (so far just
`github_release`, to avoid re-firing on the same tag) go through
`store/`, a tiny file-backed key-value store.

## Run

```
go run .
```

Reads and runs `rules.conf`. See `examples/` for a full worked
example (GitHub release → Discord alert).
