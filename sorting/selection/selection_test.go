package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	assert := assert.New(t)
	tt := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 2, 3, 4},
		},
		{
			input:    []int{4, 7, 9, 1, 2, 6, 8, 3, 5},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			input:    []int{5, 1, 2, 5, 6, 2, 2, 2, 2},
			expected: []int{1, 2, 2, 2, 2, 2, 5, 5, 6},
		},
		{
			input:    []int{4, 3, 2, 1},
			expected: []int{1, 2, 3, 4},
		},
	}
	for _, t := range tt {
		assert.Equal(t.expected, sort(t.input))
	}

}
