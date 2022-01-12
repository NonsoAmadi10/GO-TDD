package main

import (
	"testing"
)

func TestRepeat(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got , want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("It should Repeat a number based on the number of times specified", func(t *testing.T){
		got := Repeat("a",5)
		want := "aaaaa"

		assertCorrectMessage(t, got, want)
	})

	t.Run("It should not repeat a number if the repeat time is absent", func(t *testing.T) {
		got := Repeat("a", 0)
		want := "a"

		assertCorrectMessage(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a", 5)
    }
}