package test

import (
	"testing"

	"github.com/mytkdals93/go-with-me/sorting/quick"
)

func TestQuickSort(t *testing.T) {
	helper(t, quick.Sort)
}
