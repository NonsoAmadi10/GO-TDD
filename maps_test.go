package main

import "testing"

func TestSearch( t *testing.T) {
	dictionary := Dictionary{"test": "this is just an example"}

	t.Run("it should return a word from the dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		want:= "this is just an example"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
		
	})

	t.Run("it should return error message for unknown words", func(t *testing.T) {
		_, got := dictionary.Search("taofeek")
		

		assertStrings(t, got, ErrNotFound)
	})

	
}

func assertStrings(t testing.TB, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }
}

func TestAddWord(t *testing.T) {
	dictionary := Dictionary{}
	definition := "this is just a test"
	word := "test"
	dictionary.Add(word, definition)
	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
    t.Helper()

    got, err := dictionary.Search(word)
    if err != nil {
        t.Fatal("should find added word:", err)
    }

    if definition != got {
        t.Errorf("got %q want %q", got, definition)
    }
}