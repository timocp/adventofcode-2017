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

func TestTripSecurity(t *testing.T) {
	r := MustReadFirewall(bytes.NewBufferString(in)).Run().Severity
	exp := 24
	if r != exp {
		t.Errorf("Severity => %d, want %d", r, exp)
	}
}
