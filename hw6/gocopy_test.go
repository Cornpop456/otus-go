package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopy(t *testing.T) {
	type args struct {
		from   string
		to     string
		limit  int64
		offset int64
	}
	tests := []struct {
		args      args
		wantBytes int64
	}{
		{args{"files/text1.txt", "files/text2.txt", 0, 0}, 9},
		{args{"files/text1.txt", "files/text2.txt", 3, 0}, 3},
		{args{"files/text1.txt", "files/text2.txt", 0, 4}, 5},
		{args{"files/text1.txt", "files/text2.txt", 2, 7}, 2},
		{args{"files/text1.txt", "files/text2.txt", 7, 7}, 2},
	}
	for _, tc := range tests {
		err := Copy(tc.args.from, tc.args.to, tc.args.limit, tc.args.offset)

		assert.Nil(t, err)

		target, _ := os.Open(tc.args.to)

		defer target.Close()

		stat, _ := target.Stat()

		size := stat.Size()

		assert.Equal(t, tc.wantBytes, size, "bytes number should be equal")
	}
}
