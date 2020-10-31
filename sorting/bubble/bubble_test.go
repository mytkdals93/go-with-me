package bubble

import "fmt"

func ExampleBubbleSort() {
	unsorted := []int{5, 2, 1, 4, 8, 3, 7, 9, 6, 10, -1, 99, 0, 2}
	fmt.Println(Sort(unsorted))
	//Output:
	//[1 2 3 4 5 6 7 8 9]
}
