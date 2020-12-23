package main

import (
	"fmt"
	"strconv"
	"strings"
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
}

func NewRing(vals []int) *Ring {
	r := new(Ring)
	var p *Element
	for _, v := range vals {
		e := new(Element)
		if p == nil {
			r.Current = e
		} else {
			p.Next = e
		}
		e.Value = v
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
	return head, tail
}

func (e *Element) InsertAfter(head *Element, tail *Element) {
	enext := e.Next
	e.Next = head
	tail.Next = enext
}

func (r *Ring) Find(v int) *Element {
	if r.Current.Value == v {
		return r.Current
	}

	for e := r.Current.Next; e != r.Current; e = e.Next {
		if e.Value == v {
			return e
		}
	}
	return nil
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
			i = 9
		}
		e := r.Find(i)
		if e != nil {
			e.InsertAfter(head, tail)
			r.Current = r.Current.Next
			return
		}
	}
	panic(fmt.Errorf("Unable to find appropriate cup in ring"))
}

func solve(s string) string {
	var start []int

	for _, c := range s {
		d := int(c - 48)
		start = append(start, d)
	}

	r := NewRing(start)

	for i := 0; i < 100; i++ {
		step(r)
	}

	e := r.Find(1)
	r.Current = e.Next
	vals := r.GetSlice()
	var digits []string
	for _, v := range vals[:len(vals)-1] {
		digits = append(digits, strconv.Itoa(v))
	}

	return strings.Join(digits, "")
}
