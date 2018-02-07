package b_test

import (
	"testing"

	. "github.com/lukasmalkmus/issue23729/pkg/b"
)

// TestB validates that A adds two integer values correctly.
func TestB(t *testing.T) {
	if got := B("Hello", "World"); got != "Hello World" {
		t.Errorf("want: %s; got: %s", "Hello World", got)
	}
}

func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		B("Hello", "World")
	}
}
