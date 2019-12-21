package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	tests := []struct {
		i Item
	}{
		{Item{}},
		{Item{value: 10}},
		{Item{value: "Hello"}},
		{Item{value: []int{1, 2, 3}}},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.i.value, tc.i.Value(), "values should be equal")
	}
}

func TestNext(t *testing.T) {
	tests := []struct {
		i Item
	}{
		{Item{}},
		{Item{next: &Item{value: 10}}},
		{Item{next: &Item{value: "hello"}}},
	}
	for _, tc := range tests {
		assert.Same(t, tc.i.next, tc.i.Next(), "pointers should be equal")
	}
}

func TestPrev(t *testing.T) {
	tests := []struct {
		i Item
	}{
		{Item{}},
		{Item{prev: &Item{value: 20}}},
		{Item{prev: &Item{value: "world"}}},
	}
	for _, tc := range tests {
		assert.Same(t, tc.i.prev, tc.i.Prev(), "pointers should be equal")
	}
}

func TestPushFront(t *testing.T) {
	lst := &List{}

	tests := []struct {
		l          *List
		want       interface{}
		wantLength uint32
	}{
		{lst, 10, 1},
		{lst, 20, 2},
		{lst, 30, 3},
		{lst, "hello", 4},
	}
	for _, tc := range tests {
		tc.l.PushFront(tc.want)
		assert.Equal(t, tc.want, tc.l.First().Value(), "values should be equal")
		assert.Equal(t, tc.wantLength, tc.l.length, "values should be equal")
	}
}

func TestPushBack(t *testing.T) {
	lst := &List{}

	tests := []struct {
		l          *List
		want       interface{}
		wantLength uint32
	}{
		{lst, 40, 1},
		{lst, 50, 2},
		{lst, 60, 3},
		{lst, "world", 4},
	}
	for _, tc := range tests {
		tc.l.PushBack(tc.want)
		assert.Equal(t, tc.want, tc.l.Last().Value(), "values should be equal")
		assert.Equal(t, tc.wantLength, tc.l.length, "values should be equal")
	}
}

func TestLen(t *testing.T) {
	lst := &List{}

	for _, v := range []uint32{1, 2, 3, 4, 5} {
		lst.PushBack(v)
		assert.Equal(t, lst.length, lst.Len(), "values should be equal")
	}
}

func TestFirst(t *testing.T) {
	lst := &List{}

	tests := []struct {
		l        *List
		toInsert interface{}
	}{
		{lst, 0},
		{lst, 1},
		{lst, 2},
		{lst, "hello"},
	}
	for _, tc := range tests {
		tc.l.PushFront(tc.toInsert)
		assert.Same(t, tc.l.first, tc.l.First(), "pointers should be equal")
	}
}

func TestLast(t *testing.T) {
	lst := &List{}

	tests := []struct {
		l        *List
		toInsert interface{}
	}{
		{lst, 3},
		{lst, 4},
		{lst, 5},
		{lst, "world"},
	}
	for _, tc := range tests {
		tc.l.PushBack(tc.toInsert)
		assert.Same(t, tc.l.last, tc.l.Last(), "pointers should be equal")
	}
}

func TestRemove(t *testing.T) {
	lst := &List{}

	lst.PushBack(1)
	lst.PushBack(2)
	lst.PushBack(3)
	lst.PushBack(4)
	lst.PushBack(5)
	lst.PushBack(6)
	lst.PushBack(7)

	tests := []struct {
		l         *List
		wantFirst interface{}
		wantLast  interface{}
	}{
		{lst, 2, 6},
		{lst, 4, 4},
	}
	for _, tc := range tests {
		tc.l.Remove(*(tc.l.First()))
		tc.l.Remove(*(tc.l.Last()))
		tc.l.Remove(*(tc.l.Last().Prev()))
		assert.Equal(t, tc.wantFirst, tc.l.First().Value(), "values should be equal")
		assert.Equal(t, tc.wantLast, tc.l.Last().Value(), "values should be equal")
	}
}
