package test

import (
	"testing"

	"github.com/mytkdals93/go-with-me/sorting/heap"
)

func TestHeapSort(t *testing.T) {
	helper(t, heap.Sort)
}
