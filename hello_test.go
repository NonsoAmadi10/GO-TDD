package main 

import "testing"

func TestHello(t *testing.T){

	assertCorrectMessage := func(t testing.TB, got , want string){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to People", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello Chris"

		assertCorrectMessage(t, got, want)
	})
	
	t.Run("say Hello! even when no greeting is supplied in the function", func(t *testing.T) {
		got := Hello("")
		want := "Hello !"

		assertCorrectMessage(t, got, want)
	})
}