package test

import (
	"testing"

	"github.com/mytkdals93/go-with-me/sorting/merge"
)

func TestMergeSort(t *testing.T) {
	helper(t, merge.Sort)
}
