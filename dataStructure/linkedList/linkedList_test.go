package linkedList

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeLinkedList(a ...int) linkedList {
	ll := linkedList{}
	for _, v := range a {
		ll.push(&node{data: v})
	}
	return ll
}
func TestPush(t *testing.T) {
	assert := assert.New(t)
	arr := []int{10, 20, 30, 40}
	ll := linkedList{}
	for _, v := range arr {
		ll.push(&node{data: v})
	}
	assert.Equal(ll.length, len(arr))
	assert.Equal(ll.head.data, 10)
	assert.Equal(ll.tail.data, 40)

	buf := bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("10=>20=>30=>40", buf.String())
}
func TestPrepend(t *testing.T) {
	assert := assert.New(t)
	arr := []int{10, 20, 30, 40}
	ll := linkedList{}
	for _, v := range arr {
		ll.prepend(&node{data: v})
	}
	assert.Equal(ll.length, len(arr))
	assert.Equal(ll.head.data, 40)
	assert.Equal(ll.tail.data, 10)
	buf := bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("40=>30=>20=>10", buf.String())
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	//HEAD 삭제-----------------------------------------
	ll := makeLinkedList(10, 20, 30, 40, 50)
	targetData := 10
	n, err := ll.delete(targetData)
	assert.Equal(n, targetData, "check the return value")
	assert.Equal(ll.length, 4, "check length of list")
	assert.Equal(ll.head.data, 20, "check the value of head")
	assert.Equal(ll.tail.data, 50, "check the value of tail")
	buf := bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("20=>30=>40=>50", buf.String(), "check the result of print")
	assert.NoError(err)
	//중간값 삭제-----------------------------------------
	ll = makeLinkedList(10, 20, 30, 40, 50)
	targetData = 30
	n, err = ll.delete(targetData)
	assert.Equal(n, targetData, "check the return value")
	assert.Equal(ll.length, 4, "check length of list")
	assert.Equal(ll.head.data, 10, "check the value of head")
	assert.Equal(ll.tail.data, 50, "check the value of tail")
	buf = bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("10=>20=>40=>50", buf.String(), "check the result of print")
	assert.NoError(err)
	//Tail 삭제-----------------------------------------
	ll = makeLinkedList(10, 20, 30, 40, 50)
	targetData = 50
	n, err = ll.delete(targetData)
	assert.Equal(n, targetData, "check the return value")
	assert.Equal(ll.length, 4, "check length of list")
	assert.Equal(ll.head.data, 10, "check the value of head")
	assert.Equal(ll.tail.data, 40, "check the value of tail")
	buf = bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("10=>20=>30=>40", buf.String(), "check the result of print")
	assert.NoError(err)

	// 데이터가 하나 밖에 없을 경우 -----------------------
	ll = makeLinkedList(10)
	targetData = 10
	n, err = ll.delete(targetData)
	assert.Equal(n, targetData, "check the return value")
	assert.Equal(ll.length, 0, "check length of list")
	assert.Nil(ll.head)
	assert.Nil(ll.tail)
	buf = bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("", buf.String(), "check the result of print")
	assert.NoError(err)
	// Empty List에서 삭제 시도
	ll = makeLinkedList()
	n, err = ll.delete(1)
	assert.Error(err, err.Error())
	assert.Contains(err.Error(), "Empty List")
	//리스트에 존재하지 않는 값 삭제-----------------------------------------
	ll = makeLinkedList(10, 20, 30, 40, 50)
	n, err = ll.delete(1000)
	assert.Error(err, err.Error())
	assert.Contains(err.Error(), "Not Found")
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	ll := makeLinkedList(10, 20, 30, 40)

	n, err := ll.get(3)
	assert.Equal(n, ll.tail)
	assert.NoError(err)
	buf := bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("10=>20=>30=>40", buf.String())
	//index값이 음수 error처리
	n, err = ll.get(-1)
	assert.Error(err)
	//legth보다 index값이 더 큰 경우 error처리
	n, err = ll.get(ll.length + 1)
	assert.Error(err)

}
func TestPrint(t *testing.T) {
	assert := assert.New(t)
	ll := makeLinkedList(10, 20, 30, 40)

	buf := bytes.NewBufferString("")
	ll.print(buf)
	assert.Equal("10=>20=>30=>40", buf.String())

}
