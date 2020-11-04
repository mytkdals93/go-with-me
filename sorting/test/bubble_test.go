package test

import (
	"testing"

	"github.com/mytkdals93/go-with-me/sorting/bubble"
)

func TestBubbleSort(t *testing.T) {
	helper(t, bubble.Sort)
}
