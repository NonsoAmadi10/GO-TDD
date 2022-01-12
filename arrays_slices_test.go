package main

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSum(t *testing.T){

	assertCorrectMessage := func(t testing.TB, got , want int){
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", strconv.Itoa(got), strconv.Itoa(want))
		}
	}

	t.Run("It should sum slices of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4}

		got := Sum(numbers)
		want := 10 

		assertCorrectMessage(t, got, want)
	})

}

func TestSumAll( t *testing.T) {
	

	t.Run("It should sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
    	want := []int{3, 9}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}