package coding

import (
	"strconv"
)

//문제:
//https://programmers.co.kr/learn/courses/30/lessons/42587?language=go

type Printer struct {
	priority int
	location int
	target   bool
}

func (p *Printer) String() string {
	if p.target {
		return strconv.Itoa(p.priority) + "*"
	}
	return strconv.Itoa(p.priority)
}

type PrinterHeap struct {
	list []*Printer
}

func (ph *PrinterHeap) Push(p *Printer) {
	ph.list = append(ph.list, p)
}
func (ph *PrinterHeap) Pop() *Printer {
outter:
	for {
		top := ph.list[0]
		for _, p := range ph.list {
			if p.priority > top.priority {
				ph.list = ph.list[1:]
				ph.list = append(ph.list, top)
				continue outter
			}
		}
		break
	}
	top := ph.list[0]
	ph.list = ph.list[1:]
	return top
}

func solution1(priorities []int, location int) int {
	ph := PrinterHeap{}
	for i, p := range priorities {
		if i == location {
			ph.Push(&Printer{
				priority: p,
				location: i,
				target:   true,
			})
		} else {
			ph.Push(&Printer{
				priority: p,
				location: i,
			})
		}

	}
	count := 0
	for len(ph.list) > 0 {
		count++
		p := ph.Pop()
		if p.target {
			break
		}
	}

	return count
}
