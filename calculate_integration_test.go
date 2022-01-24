// go:build integration
package main 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {

	t.Run("it should increase a number by 2", func(t *testing.T) {
		got := Calculate(2)
		want := 4

		assert.Equal(t, got, want)
		

	})
	

}