package iteration

import "testing"

func TestIteration(t *testing.T) {
	got := Iteration("a", 6)
	want := "aaaaaa"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func BenchmarkIteration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Iteration("a", 6)
	}
}
