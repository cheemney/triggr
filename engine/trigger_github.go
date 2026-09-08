package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/cheemney/triggr/store"
)

// GitHubReleaseTrigger polls a repo's latest release endpoint and
// fires only when the tag differs from the last one it saw, tracked
// through Store so it survives restarts.
type GitHubReleaseTrigger struct {
	Owner    string
	Repo     string
	Interval time.Duration
	Store    *store.Store
}

type githubRelease struct {
	TagName string `json:"tag_name"`
}

func (t *GitHubReleaseTrigger) Watch() error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", t.Owner, t.Repo)
	key := fmt.Sprintf("github_release:%s/%s", t.Owner, t.Repo)

	for {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}

		var release githubRelease
		err = json.NewDecoder(resp.Body).Decode(&release)
		resp.Body.Close()
		if err != nil {
			return err
		}

		if release.TagName != "" {
			if last, seen := t.Store.Get(key); !seen || last != release.TagName {
				if err := t.Store.Set(key, release.TagName); err != nil {
					return err
				}
				return nil
			}
		}

		time.Sleep(t.Interval)
	}
}
