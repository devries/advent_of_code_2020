package main

import (
	"fmt"
)

var input = "476138259"
var testInput = "389125467"

func main() {
	r := solve(input)
	fmt.Println(r)
}

type Element struct {
	Value int
	Next  *Element
}

type Ring struct {
	Current *Element
	Locator map[int]*Element
}

func NewRing(vals []int) *Ring {
	r := new(Ring)
	r.Locator = make(map[int]*Element)

	var p *Element
	for _, v := range vals {
		e := new(Element)
		if p == nil {
			r.Current = e
		} else {
			p.Next = e
		}
		e.Value = v
		r.Locator[v] = e
		p = e
	}

	for i := 10; i <= 1000000; i++ {
		e := new(Element)
		e.Value = i
		r.Locator[i] = e
		p.Next = e
		p = e
	}

	p.Next = r.Current

	return r
}

// Cut n items from ring and return head and tail
func (r *Ring) Cut(n int) (*Element, *Element) {
	head := r.Current.Next
	tail := r.Current
	for i := 0; i < n; i++ {
		tail = tail.Next
	}

	r.Current.Next = tail.Next
	tail.Next = nil
	return head, tail
}

func (e *Element) InsertAfter(head *Element, tail *Element) {
	enext := e.Next
	e.Next = head
	tail.Next = enext
}

func (r *Ring) Find(v int, cut *Element) *Element {
	/*
		if r.Current.Value == v {
			return r.Current
		}

		for e := r.Current.Next; e != r.Current; e = e.Next {
			if e.Value == v {
				return e
			}
		}
		return nil
	*/
	for cutE := cut; cutE != nil; cutE = cutE.Next {
		if cutE.Value == v {
			return nil
		}
	}
	e := r.Locator[v]
	return e
}

func (r *Ring) GetSlice() []int {
	s := []int{r.Current.Value}

	for e := r.Current.Next; e != r.Current; e = e.Next {
		s = append(s, e.Value)
	}

	return s
}

func step(r *Ring) {
	current := r.Current
	head, tail := r.Cut(3)
	cv := current.Value

	for i := cv - 1; i != cv; i-- {
		if i < 0 {
			i = 1000000
		}
		e := r.Find(i, head)
		if e != nil {
			e.InsertAfter(head, tail)
			r.Current = r.Current.Next
			return
		}
	}
	panic(fmt.Errorf("Unable to find appropriate cup in ring"))
}

func solve(s string) int64 {
	var start []int

	for _, c := range s {
		d := int(c - 48)
		start = append(start, d)
	}

	r := NewRing(start)

	for i := 0; i < 10000000; i++ {
		step(r)
	}

	var result int64 = 1
	e := r.Locator[1]
	for i := 0; i < 3; i++ {
		result = result * int64(e.Value)
		e = e.Next
	}

	return result
}
