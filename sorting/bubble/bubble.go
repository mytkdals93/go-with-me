package bubble

func Sort(arr []int) []int {
	return bubbleSort1(arr)
}

func bubbleSort1(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := len(arr) - 1; j > i; j-- {
			//Copmare two numbers at the end of the sequence
			if arr[j-1] > arr[j] {
				//if the number on the left is bigger than the number on the right,
				//the numbers will be swapped.
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
		//In one round of operations, the smllest value in the sequense has moved to the left edge.
		//Which means [:i] is sorted. So, we just going to compare until 'i' in the second iteration
	}
	return arr
}

func bubbleSort2(arr []int) []int {
	//swapped flag is using to found out the sequnce has swapped.
	//if the loop ends and swapped is still equaql to false, our algorithm will assume the list is fully sorted.
	swapped := true
	unsorted := len(arr)
	for swapped {
		swapped = false
		for first, second := 0, 1; second < unsorted; first, second = first+1, second+1 {
			if arr[first] > arr[second] {
				arr[first], arr[second] = arr[second], arr[first]
				swapped = true
			}
		}
		unsorted--
	}
	return arr
}
