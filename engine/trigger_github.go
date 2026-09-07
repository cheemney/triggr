package engine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// GitHubReleaseTrigger polls a repo's latest release endpoint and
// fires whenever it gets a response with a tag. No de-duplication —
// it'll fire again on the same release next poll. Known gap, tracked
// separately rather than patched in here.
type GitHubReleaseTrigger struct {
	Owner    string
	Repo     string
	Interval time.Duration
}

type githubRelease struct {
	TagName string `json:"tag_name"`
}

func (t *GitHubReleaseTrigger) Watch() error {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", t.Owner, t.Repo)

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
			return nil
		}

		time.Sleep(t.Interval)
	}
}
