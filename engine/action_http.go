package engine

import (
	"fmt"
	"net/http"
)

// HTTPAction fires a request to URL. Method defaults to GET if empty.
// No body or custom headers yet — that's more than this phase needs.
type HTTPAction struct {
	Method string
	URL    string
}

func (a *HTTPAction) Execute() error {
	method := a.Method
	if method == "" {
		method = http.MethodGet
	}

	req, err := http.NewRequest(method, a.URL, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Printf("http action: %s %s -> %s\n", method, a.URL, resp.Status)
	return nil
}
