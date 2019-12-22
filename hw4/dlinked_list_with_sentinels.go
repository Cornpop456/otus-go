package main

import "fmt"

// List2 is realization of doubly linked list
type List2 struct {
	firstFake *Item2
	lastFake  *Item2
	length    uint32
}

// Item2 of List2
type Item2 struct {
	value interface{}
	next  *Item2
	prev  *Item2
}

// Value returns value of Item2
func (i Item2) Value() interface{} {
	return i.value
}

// Next returns next Item2 of List2
func (i Item2) Next() *Item2 {
	return i.next
}

// Prev returns previous Item2 of List2
func (i Item2) Prev() *Item2 {
	return i.prev
}

// Len returns length of List2
func (l *List2) Len() uint32 {
	return l.length
}

// First returns first Item2 of List2
func (l *List2) First() *Item2 {
	return l.firstFake.Next()
}

// Last returns last Item2 of List2
func (l *List2) Last() *Item2 {
	return l.lastFake.Prev()
}

// PushFront pushes value in start of List2
func (l *List2) PushFront(v interface{}) {
	newItem := &Item2{value: v}
	fst := l.firstFake.Next()
	newItem.prev = l.firstFake
	newItem.next = fst
	fst.prev = newItem
	l.firstFake.next = newItem
	l.length++
}

// PushBack pushes value in end of List2
func (l *List2) PushBack(v interface{}) {
	newItem := &Item2{value: v}
	last := l.lastFake.Prev()
	newItem.prev = last
	newItem.next = l.lastFake
	last.next = newItem
	l.lastFake.prev = newItem
	l.length++
}

// Remove deletes Item2 from List2
func (l *List2) Remove(i Item2) {
	prev, next := i.Prev(), i.Next()
	prev.next, next.prev = next, prev
	l.length--
}

// Traverse prints elements of List2 from first to last
func (l *List2) Traverse() {
	cur := l.firstFake.Next()
	if cur == nil {
		fmt.Printf("empty list")
	}
	for cur != l.lastFake {
		if cur.Next() == l.lastFake {
			fmt.Printf("%v", cur.Value())
			fmt.Println()
			return
		}
		fmt.Printf("%v -> ", cur.Value())
		cur = cur.Next()
	}
}

// TraverseBack prints elements of List2 from last to first
func (l *List2) TraverseBack() {
	cur := l.lastFake.Prev()
	if cur == nil {
		fmt.Printf("empty list")
	}
	for cur != l.firstFake {
		if cur.Prev() == l.firstFake {
			fmt.Printf("%v", cur.Value())
			fmt.Println()
			return
		}
		fmt.Printf("%v <- ", cur.Value())
		cur = cur.Prev()
	}
}

func main() {
	firstFake := &Item2{}
	lastFake := &Item2{}
	firstFake.next = lastFake
	lastFake.prev = firstFake
	list := &List2{firstFake: firstFake, lastFake: lastFake}
	list.PushBack(1)
	list.PushFront(11)
	list.PushBack(2)
	list.PushBack(3)
	list.PushBack(4)
	list.PushBack(5)
	list.PushBack(6)
	list.Traverse()
	list.TraverseBack()
	list.Remove(*(list.First().Next().Next()))
	list.Traverse()
	list.TraverseBack()
	list.Remove(*(list.Last().Prev()))
	list.Traverse()
	list.TraverseBack()
}
