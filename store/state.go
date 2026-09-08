package store

import (
	"bufio"
	"os"
	"strings"
)

// Store is a tiny file-backed key-value store, one key=value pair
// per line. No locking — fine for a single process, not safe for
// concurrent writers.
type Store struct {
	path string
}

func New(path string) *Store {
	return &Store{path: path}
}

func (s *Store) Get(key string) (string, bool) {
	f, err := os.Open(s.path)
	if err != nil {
		return "", false
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		k, v, ok := strings.Cut(scanner.Text(), "=")
		if ok && k == key {
			return v, true
		}
	}
	return "", false
}

func (s *Store) Set(key, value string) error {
	entries := map[string]string{}

	if f, err := os.Open(s.path); err == nil {
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			if k, v, ok := strings.Cut(scanner.Text(), "="); ok {
				entries[k] = v
			}
		}
		f.Close()
	}

	entries[key] = value

	f, err := os.Create(s.path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for k, v := range entries {
		if _, err := w.WriteString(k + "=" + v + "\n"); err != nil {
			return err
		}
	}
	return w.Flush()
}
