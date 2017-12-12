package day12

import (
	"bytes"
	"testing"
)

var testProgram = `0 <-> 2
1 <-> 1
2 <-> 0, 3, 4
3 <-> 2, 4
4 <-> 2, 3, 6
5 <-> 6
6 <-> 4, 5`

func TestCountConnected(t *testing.T) {
	r := MustReadProgramList(bytes.NewBufferString(testProgram)).CountConnected(0)
	if r != 6 {
		t.Errorf("CountConnected => %d, want %d", r, 6)
	}
}

func TestCountGroups(t *testing.T) {
	r := MustReadProgramList(bytes.NewBufferString(testProgram)).CountGroups()
	if r != 2 {
		t.Errorf("CountGroups => %d, want %d", r, 2)
	}
}
