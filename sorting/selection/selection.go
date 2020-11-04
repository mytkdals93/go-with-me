package selection

func Sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		//intialize maxIndex to the first index of Unsorted part
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			//find smallest element and
			//put the index of smallest element into minIndex
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		//swap the frist index of unsorted part and the smallest element
		//In one round of operations, the smllest value in the sequense has moved to the next index of sorted part.
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}
