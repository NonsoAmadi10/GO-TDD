package main 

import "testing"

func TestHello(t *testing.T){
	got := Hello()
	want := "Hello World"

	if got != want {
		t.Errorf("got %q instead of %q", got, want)
	}
}