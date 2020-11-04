package merge

func Sort(arr []int) []int {
	return merge_sort(arr)
}

func merge_sort(arr []int) []int {
	middle := len(arr) / 2

	if len(arr) == 1 {
		return arr
	}
	left := make([]int, middle)
	right := make([]int, len(arr)-middle)

	for i := 0; i < len(arr); i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}

	}

	return merge(merge_sort(left), merge_sort(right))
}

func merge(left, right []int) []int {
	result := make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return result
}
