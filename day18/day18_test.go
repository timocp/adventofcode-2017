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
	d := NewDuet()
	err := d.Load(bytes.NewBufferString(in))
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

var in2 = `snd 1
snd 2
snd p
rcv a
rcv b
rcv c
rcv d
`

func TestPart2(t *testing.T) {
	r := Part2(in2)
	if r != 3 {
		t.Errorf("Part2() => %d, want %d", r, 3)
	}
}
