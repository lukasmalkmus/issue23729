package a_test

import (
	"testing"

	. "github.com/lukasmalkmus/issue23729/pkg/a"
)

// TestA validates that A adds two integer values correctly.
func TestA(t *testing.T) {
	if got := A(1, 2); got != 3 {
		t.Errorf("want: %d; got: %d", 3, got)
	}
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		A(1, 2)
	}
}
