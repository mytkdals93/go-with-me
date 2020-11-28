package merge

func Sort(arr []int) []int {
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}
	middle := len(arr) / 2

	left := arr[:middle]
	right := arr[middle:]

	return merge(mergeSort(left), mergeSort(right))
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
	for len(left) > 0 {
		result[i] = left[0]
		left = left[1:]
		i++
	}
	for len(right) > 0 {
		result[i] = right[0]
		right = right[1:]
		i++
	}

	return result
}
