package linkedList

import (
	"fmt"
	"io"
	"os"
)

type node struct {
	data int
	next *node
}
type linkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedList) push(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
		return
	}
	l.tail.next = n
	l.tail = n
	l.length++

}

func (l *linkedList) prepend(n *node) {
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
		return
	}
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}
func (l *linkedList) delete(value int) (int, error) {
	//Empty List
	if l.head == nil {
		err := fmt.Errorf("Empty List")
		return 0, err
	}

	//지우고자 하는 값이 HEAD인 경우
	if l.head.data == value {
		// HEAD와 TAIL이 같은 경우
		if l.head == l.tail {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
		}
		l.length--
		return value, nil
	}

	temp := l.head
	for temp.next != nil && temp.next.data != value {
		temp = temp.next
	}

	//지우고자 하는 값이 LIST에 없는 경우
	if temp.next == nil {
		err := fmt.Errorf("Not Found")
		return 0, err
	}

	//지우고자 하는 값이 TAIL인 경우
	if temp.next == l.tail {
		l.tail = temp
		l.tail.next = nil
	} else {
		temp.next = temp.next.next
	}
	l.length--
	return value, nil
}
func (l *linkedList) get(index int) (*node, error) {
	if index < 0 {
		err := fmt.Errorf("invalid index: index shouldn't be negative")
		return nil, err
	}
	if !(index < l.length) {
		err := fmt.Errorf("Out of Length,length:%d index:%d", l.length, index)
		return nil, err
	}
	temp := l.head
	for index > 0 {
		temp = temp.next
		index--
	}

	return temp, nil
}
func (l linkedList) print(w io.Writer) {
	if w == nil {
		w = os.Stdout
	}
	temp := l.head

	for i := 0; i < l.length; i++ {
		if i != l.length-1 {
			fmt.Fprint(w, temp.data, "=>")
		} else {
			fmt.Fprint(w, temp.data)
		}
		temp = temp.next
	}
}
