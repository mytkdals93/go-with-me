package quick

import (
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Sort(arr []int) []int {
	return quickSort(arr)
}

func quickSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	p := r.Intn(len(arr))
	arr[p], arr[len(arr)-1] = arr[len(arr)-1], arr[p]
	p = len(arr) - 1
	i := 0
	j := p - 1

	for i <= j {
		if arr[i] < arr[p] {
			i++
		} else {
			if arr[j] >= arr[p] {
				j--
			} else {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
	}
	arr[p], arr[i] = arr[i], arr[p]

	quickSort(arr[:i])
	quickSort(arr[i+1:])
	return arr
}
