package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	val1, _ := Unpack(`qwe\4\5d`)
	val2, _ := Unpack(`qwe\45`)
	val3, _ := Unpack(`qwe\\5`)
	val4, _ := Unpack("a4бc2d5eз5")
	val5, _ := Unpack("abcd")
	val6, _ := Unpack(`a4\43`)

	assert.Equal(t, `qwe45d`, val1, "They should be equal")
	assert.Equal(t, `qwe44444`, val2, "They should be equal")
	assert.Equal(t, `qwe\\\\\`, val3, "They should be equal")
	assert.Equal(t, "aaaaбccdddddeззззз", val4, "They should be equal")
	assert.Equal(t, "abcd", val5, "They should be equal")
	assert.Equal(t, `aaaa444`, val6, "They should be equal")

	_, err1 := Unpack("45")
	_, err2 := Unpack(`b4\a`)
	_, err3 := Unpack(`b445\\`)
	_, err4 := Unpack(`a3b6ff\\\`)

	assert.NotNil(t, err1)
	assert.NotNil(t, err2)
	assert.NotNil(t, err3)
	assert.NotNil(t, err4)
}
