package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	tests := []struct {
		i    Item
		want interface{}
	}{
		{Item{}, nil},
		{Item{value: 10}, 10},
		{Item{value: "Hello"}, "Hello"},
		{Item{value: []int{1, 2, 3}}, []int{1, 2, 3}},
	}
	for _, tc := range tests {
		assert.Equal(t, tc.want, tc.i.Value(), "values should be equal")
	}
}

func TestNext(t *testing.T) {
	items := []*Item{
		&Item{value: 10},
		&Item{value: "hello"},
	}
	tests := []struct {
		i    Item
		want *Item
	}{
		{Item{}, nil},
		{Item{next: items[0]}, items[0]},
		{Item{next: items[1]}, items[1]},
	}
	for _, tc := range tests {
		assert.Same(t, tc.want, tc.i.Next(), "pointers should be equal")
	}
}

func TestPrev(t *testing.T) {
	items := []*Item{
		&Item{value: 20},
		&Item{value: "world"},
	}
	tests := []struct {
		i    Item
		want *Item
	}{
		{Item{}, nil},
		{Item{prev: items[0]}, items[0]},
		{Item{prev: items[1]}, items[1]},
	}
	for _, tc := range tests {
		assert.Same(t, tc.want, tc.i.Prev(), "pointers should be equal")
	}
}
