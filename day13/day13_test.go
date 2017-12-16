package day13

import (
	"bytes"
	"testing"
)

var in = `0: 3
1: 2
4: 4
6: 4
`

func TestSeverity(t *testing.T) {
	r := MustReadFirewall(bytes.NewBufferString(in)).Run().Severity
	exp := 24
	if r != exp {
		t.Errorf("Severity => %d, want %d", r, exp)
	}
}

func TestSneakyWaitTime(t *testing.T) {
	r := MustReadFirewall(bytes.NewBufferString(in)).SneakyWaitTime()
	exp := 10
	if r != exp {
		t.Errorf("SneakyWaitTime() => %d, want %d", r, exp)
	}
}
