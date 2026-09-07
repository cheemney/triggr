# github-to-discord

Watches a GitHub repo's latest release and posts an alert to a
Discord channel through a webhook. No API keys needed on either side.

## Setup

1. In Discord: channel settings → Integrations → Webhooks → New
   Webhook → copy the URL.
2. Paste that URL into `rules.conf` in place of
   `https://discord.com/api/webhooks/REPLACE_ME`.
3. Swap `"golang" "go"` in `rules.conf` for whichever repo you want
   to watch.

## Run

From the repo root:

```
go run ./examples/github-to-discord
```

## Known issue

Every poll currently re-fires on the same release — there's no
de-duplication yet, so you'll get repeated Discord messages rather
than one per new release. Being tracked and fixed next.
