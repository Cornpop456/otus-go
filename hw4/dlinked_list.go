package main

import "fmt"

// List is realization of doubly linked list
type List struct {
	first  *Item
	last   *Item
	length uint32
}

// Item of List
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

// Value returns value of Item
func (i Item) Value() interface{} {
	return i.value
}

// Next returns next Item of List
func (i Item) Next() *Item {
	return i.next
}

// Prev returns previous Item of List
func (i Item) Prev() *Item {
	return i.prev
}

// Len returns length of List
func (l *List) Len() uint32 {
	return l.length
}

// First returns first Item of List
func (l *List) First() *Item {
	return l.first
}

// Last returns last Item of List
func (l *List) Last() *Item {
	return l.last
}

// PushFront pushes value in start of List
func (l *List) PushFront(v interface{}) {
	newItem := &Item{value: v}
	if l.Len() == 0 {
		l.first, l.last = newItem, newItem
	} else {
		newItem.next = l.first
		l.first.prev = newItem
		l.first = newItem
	}
	l.length++
}

// PushBack pushes value in end of List
func (l *List) PushBack(v interface{}) {
	newItem := &Item{value: v}
	if l.Len() == 0 {
		l.first, l.last = newItem, newItem
	} else {
		newItem.prev = l.last
		l.last.next = newItem
		l.last = newItem
	}
	l.length++
}

// Remove deletes Item from List
func (l *List) Remove(i Item) {
	l.length--
	prev, next := i.Prev(), i.Next()
	if next == nil && prev == nil {
		l.first = nil
		l.last = nil
		return
	} else if prev == nil {
		l.first = next
		l.first.prev = nil
	} else if next == nil {
		l.last = prev
		l.last.next = nil
	} else {
		prev.next, next.prev = next, prev
	}
}

// Traverse prints elements of List from first to last
func (l *List) Traverse() {
	cur := l.first
	if cur == nil {
		fmt.Printf("empty list")
	}
	for cur != nil {
		if cur.Next() == nil {
			fmt.Printf("%v", cur.Value())
			fmt.Println()
			return
		}
		fmt.Printf("%v -> ", cur.Value())
		cur = cur.Next()
	}
}

// TraverseBack prints elements of List from last to first
func (l *List) TraverseBack() {
	cur := l.last
	if cur == nil {
	}
	for cur != nil {
		if cur.Prev() == nil {
			fmt.Printf("%v", cur.Value())
			fmt.Println()
			return
		}
		fmt.Printf("%v <- ", cur.Value())
		cur = cur.Prev()
	}
}

// func main() {
// 	list := &List{}
// 	list.PushBack(1)
// 	list.PushBack(2)
// 	list.PushBack(3)
// 	list.PushBack(4)
// 	list.PushBack(5)
// 	list.PushBack(6)
// 	list.Traverse()
// 	list.TraverseBack()
// 	list.Remove(*(list.First().Next().Next()))
// 	list.Traverse()
// 	list.TraverseBack()
// 	list.Remove(*(list.Last().Prev()))
// 	list.Traverse()
// 	list.TraverseBack()
// }
