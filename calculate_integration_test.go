// go:build integration
package main 

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCalculate(t *testing.T) {

	assert := assert.New(t)

	t.Run("it should increase a number by 2", func(t *testing.T) {
		var tests = []struct {
			input int
			expected int 
		}{
			{2, 4},
			{-1, 1},
			{0, 2},
			{100, 102},
			{-5, -3},
		}
		
		for _, test := range tests {
			assert.Equal(Calculate(test.input), test.expected)
		}

	})
	

}