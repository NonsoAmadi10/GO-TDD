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
	
	checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }

	t.Run("It should sum all slices", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
    	want := []int{3, 9}
		
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }
	t.Run("It should sum all tails of a slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
    	want := []int{2, 9}

    	checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		checkSums := func(t testing.TB, got, want []int) {
			t.Helper()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got %v want %v", got, want)
			}
		}
        got := SumAllTails([]int{}, []int{3, 4, 5})
        want := []int{0, 9}

        checkSums(t, got, want)
    })
    
}