package day24

import (
	"bytes"
	"testing"
)

var testIn = `0/2
2/2
2/3
3/4
3/5
0/1
10/1
9/10
`

func TestMaxStrength(t *testing.T) {
	b := NewBridge(bytes.NewBufferString(testIn))
	r := b.MaxStrength()
	exp := 31
	if r != exp {
		t.Errorf("maxStrength() => %d, want %d", r, exp)
	}
}

func TestLongestBridgeStrength(t *testing.T) {
	b := NewBridge(bytes.NewBufferString(testIn))
	r := b.LongestBridgeStrength()
	exp := 19
	if r != exp {
		t.Errorf("maxStrength() => %d, want %d", r, exp)
	}
}
