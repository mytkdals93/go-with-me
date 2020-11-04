package test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type asc []int

func (a asc) Len() int           { return len(a) }
func (a asc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a asc) Less(i, j int) bool { return a[i] < a[j] }

func helper(t *testing.T, f func([]int) []int) {
	rand.Seed(int64((time.Now()).Nanosecond()))

	assert := assert.New(t)
	var tcs [][]int

	for i := 0; i < 10; i++ {
		tcs = append(tcs, []int{})
		for j := 0; j < 10; j++ {
			tcs[i] = append(tcs[i], int(rand.Uint32()%100))
		}
	}

	for _, tc := range tcs {
		actual := append([]int{}, tc...)
		sort.Sort(asc(actual))
		assert.Equal(actual, f(tc), tc)
	}
}
