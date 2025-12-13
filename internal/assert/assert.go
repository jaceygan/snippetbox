package assert

import (
	"strings"
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper()
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

func StringContains(t *testing.T, actual, substring string) {
	t.Helper()
	if !strings.Contains(actual, substring) {
		t.Errorf("expected %q to contain %q", actual, substring)
	}
}
