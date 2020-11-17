package linkedList

import (
	"fmt"
	"io"
	"os"
)

type linkedList struct {
	head   *node
	length int
}
type node struct {
	data int
	next *node
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l *linkedList) deleteWithValue(value int) (int, error) {
	if value == l.head.data {
		l.head = l.head.next
		l.length--
		return value, nil
	}
	prev := l.head
	for prev.next != nil && prev.next.data != value {
		prev = prev.next
	}
	if prev.next == nil {
		err := fmt.Errorf("Not Found")
		return 0, err
	}
	if prev.next.next != nil {
		prev.next = prev.next.next
	} else {
		prev.next = nil
	}
	l.length--
	return value, nil
}

func (l linkedList) print(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}
	temp := l.head

	for temp != nil {
		fmt.Fprint(w, temp.data, "=>")
		temp = temp.next
	}
}
