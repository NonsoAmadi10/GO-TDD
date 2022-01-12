package main

import "testing"
func TestAdd(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got , want int){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("It should add two numbers", func(t *testing.T){
		got := Add(3, 4)
		want := 7 

		assertCorrectMessage(t, got, want)
	})
}