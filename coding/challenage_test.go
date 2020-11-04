package coding

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution(t *testing.T) {
	assert := assert.New(t)
	var tcs = []struct {
		priorities []int
		location   int
		expected   int
	}{
		// {
		// 	priorities: []int{2, 1, 3, 2},
		// 	location:   2,
		// 	expected:   1,
		// },
		{
			priorities: []int{1, 1, 9, 1, 1, 1},
			location:   0,
			expected:   5,
		},
		{
			priorities: []int{2, 1, 2, 1, 3, 2},
			location:   0,
			expected:   3,
		},
		{
			priorities: []int{1, 2, 3},
			location:   0,
			expected:   3,
		},
		{
			priorities: []int{5, 4, 3, 2, 1},
			location:   3,
			expected:   4,
		},
		{
			priorities: []int{10},
			location:   0,
			expected:   1,
		},
		{
			priorities: []int{9, 9, 9, 9, 9, 9, 5, 2, 3, 5, 2, 3, 5, 5, 3, 2, 5, 5},
			location:   2,
			expected:   3,
		},
	}

	for _, tc := range tcs {
		assert.Equal(tc.expected, solution1(tc.priorities, tc.location), tc.priorities)

	}
}
