package heap

type heap []int

func (h *heap) push(val int) {
	*h = append(*h, val)

	idx := len(*h) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if (*h)[parentIdx] > (*h)[idx] {
			(*h)[parentIdx], (*h)[idx] = (*h)[idx], (*h)[parentIdx]
		}
		idx = parentIdx
	}

}

func (h *heap) pop() int {
	if len(*h) == 0 {
		return -1
	}
	top := (*h)[0]
	last := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	if len(*h) == 0 {
		return top
	}

	(*h)[0] = last
	idx := 0

	for idx < len(*h) {

		swapIdx := -1
		leftIdx := idx*2 + 1
		rightIdx := idx*2 + 2
		if leftIdx < len(*h) {
			if (*h)[leftIdx] < (*h)[idx] {
				swapIdx = leftIdx
			}
		}
		if rightIdx < len(*h) {
			if swapIdx > 0 {
				if (*h)[rightIdx] < (*h)[leftIdx] {
					swapIdx = rightIdx
				}
			} else if (*h)[rightIdx] < (*h)[idx] {
				swapIdx = rightIdx
			}
		}

		if swapIdx < 0 {
			break
		}
		(*h)[swapIdx], (*h)[idx] = (*h)[idx], (*h)[swapIdx]
		idx = swapIdx
	}

	return top
}

func Sort(arr []int) []int {
	h := &heap{}
	for _, v := range arr {
		h.push(v)
	}
	for i := range arr {
		arr[i] = h.pop()
	}
	return arr
}
