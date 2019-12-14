package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testData = []struct {
	in  string
	out []string
}{
	{`a`, []string{"a"}},
	{`a a b b b`, []string{"b", "a"}},
	{`    -    -!!! ;;;; ....,,, ,,,!
	
	`, []string{""}},
	{``, []string{""}},
	{` A A A a b B!!! a a C c C c C!!`, []string{"a", "c", "b"}},
	{`---   hello HELLO, heLl"o"!!!`, []string{"hello"}},
	{`
	a - a. a, A--'' "a" a a''' a .a. a!
	
	J:!.

	b b b!    b... b b b B,. b;;; 

	d D d d d d d 

	c C !c. ,C, c: 	C c c c C !c. ,C, c: 	C c c 
	e e: 		e.. E e e
	
	---  ...  -- - -  - 
	
	f "F" "f" f f 
	g g. 	g G!
	h h h 
	i i
	`, []string{"c", "a", "b", "d", "e", "f", "g", "h", "i", "j"}},
}

func TestTop10(t *testing.T) {
	for _, testCase := range testData {

		res := Top10(testCase.in)

		assert.Equal(t, len(testCase.out), len(res), "Expected to be equal lengths in %v %v", testCase.out, res)

		for i, v := range res {
			assert.Equal(t, testCase.out[i], v, "Expected to be equal")
		}
	}
}
