package linkedList

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeLinkedList(a ...int) linkedList {
	mylist := linkedList{}
	for _, v := range a {
		mylist.prepend(&node{data: v})
	}
	return mylist
}
func TestPrepend(t *testing.T) {
	assert := assert.New(t)

	mylist := makeLinkedList(40, 10, 20, 30)
	assert.Equal(mylist.length, 4)
	assert.Equal(mylist.head.data, 30)
	assert.Equal(mylist.head.next.data, 20)
	assert.Equal(mylist.head.next.next.data, 10)
	assert.Equal(mylist.head.next.next.next.data, 40)
}

func TestDeleteWithValue(t *testing.T) {
	assert := assert.New(t)
	mylist := makeLinkedList(40, 10, 20, 30)
	//-----------------------------------------
	originLen := mylist.length
	targetData := 40
	n, err := mylist.deleteWithValue(targetData)
	assert.Equal(n, targetData, "check the return value")

	assert.Equal(mylist.length, originLen-1, "check length of list")
	buf := bytes.NewBufferString("")
	mylist.print(buf)
	assert.Equal("30=>20=>40=>", buf.String(), "check the result of print")
	assert.NoError(err)

	n, err = mylist.deleteWithValue(1000)
	assert.NoError(err, err.Error())
}
func TestPrint(t *testing.T) {
	assert := assert.New(t)
	mylist := makeLinkedList(40, 10, 20, 30)

	buf := bytes.NewBufferString("")
	mylist.print(buf)
	assert.Equal("30=>20=>10=>40=>", buf.String())

}
