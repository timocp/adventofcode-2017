package day18

import (
	"bytes"
	"testing"
)

var in = `set a 1
add a 2
mul a a
mod a 5
snd a
set a 0
rcv a
jgz a -1
set a 1
jgz a -2`

func TestLoad(t *testing.T) {
	d := newDuet()
	err := d.load(bytes.NewBufferString(in))
	if err != nil {
		t.Errorf("load() error => %v, want nil", err)
	}
}

func TestPart1(t *testing.T) {
	r := Part1(bytes.NewBufferString(in))
	if r != 4 {
		t.Errorf("Part1() => %d, want %d", r, 4)
	}
}
