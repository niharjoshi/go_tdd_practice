package helloworld

import "testing"

func assertTestResult(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestHello(t *testing.T) {
	t.Run("Saying hello by name", func(t *testing.T) {
		got := Hello("Nihar", "")
		want := "Hello, Nihar"
		assertTestResult(t, got, want)
	})
	t.Run("Saying hello without a name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertTestResult(t, got, want)
	})
	t.Run("Saying hello in Spanish", func(t *testing.T) {
		got := Hello("Nihar", "Spanish")
		want := "Hola, Nihar"
		assertTestResult(t, got, want)
	})
	t.Run("Saying hello in French", func(t *testing.T) {
		got := Hello("Nihar", "French")
		want := "Bonjour, Nihar"
		assertTestResult(t, got, want)
	})
}
