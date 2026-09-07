package engine

import "net/http"

// WebhookTrigger spins up a tiny HTTP server and blocks until exactly
// one request hits Path, then shuts the server down. Fires once —
// a real version would need to keep listening for repeat triggers,
// that's a later problem (probably tied into the engine loop, not here).
type WebhookTrigger struct {
	Addr string // e.g. ":8080"
	Path string // e.g. "/fire"
}

func (t *WebhookTrigger) Watch() error {
	fired := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc(t.Path, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		close(fired)
	})

	server := &http.Server{Addr: t.Addr, Handler: mux}
	go server.ListenAndServe()
	defer server.Close()

	<-fired
	return nil
}
