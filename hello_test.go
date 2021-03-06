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
		got := Hello("Chris", "english")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})
	
	t.Run("say Hello! even when no greeting is supplied in the function", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in spanish", func(t *testing.T){
		got := Hello("Alonso", "spanish")
		want := "Hola, Alonso"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T){
		got := Hello("Pierre", "french")
		want := "Bonjour, Pierre"

		assertCorrectMessage(t, got, want)
	})
}