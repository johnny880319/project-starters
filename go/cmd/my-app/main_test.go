package main

import "testing"

func TestRun(t *testing.T) {
	t.Parallel()
	if err := run(); err != nil {
		t.Errorf("run() returned an error: %v", err)
	}
}
