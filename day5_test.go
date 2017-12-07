package main

import (
	"bytes"
	"testing"
)

func TestString(t *testing.T) {
	tramp := sample()
	r := tramp.String()
	exp := "(0) 3  0  1  -3 "
	if r != exp {
		t.Errorf("tramp.String() => %s, want %s", r, exp)
	}
}

func TestJump(t *testing.T) {
	tramp := sample()
	for i, exp := range []string{
		"(1) 3  0  1  -3 ",
		" 2 (3) 0  1  -3 ",
		" 2  4  0  1 (-3)",
		" 2 (4) 0  1  -2 ",
		" 2  5  0  1  -2 ",
	} {
		if tramp.Escaped() {
			t.Errorf("%dth tramp.Escaped() => true, want false", i+1)
		}
		tramp.Jump()
		if tramp.String() != exp {
			t.Errorf("%dth tramp.Jump() => %s, want %s", i+1, tramp.String(), exp)
			break
		}
	}
	// last jump should escape
	if !tramp.Escaped() {
		t.Errorf("Final tramp.Escaped() => false, want true")
	}
}

func TestStepsToExit(t *testing.T) {
	exp := 5
	count := sample().StepsToExit()
	if count != exp {
		t.Errorf("StepsToExit(%v) => %d, want %d", sample(), count, exp)
	}
}

func TestNewReadTrampoline(t *testing.T) {
	input := "11\n22\n33\n44\n55\n"
	exp := "(11) 22  33  44  55 "
	tramp, err := NewReadTrampoline(bytes.NewBufferString(input))
	if err != nil {
		t.Errorf("NewRead(%s) returned error %v, expected nil", input, err)
	} else if tramp.String() != exp {
		t.Errorf("NewRead(%s) => %s, want %s", input, tramp.String(), exp)
	}
}

func sample() *Trampoline {
	return NewTrampoline([]int{0, 3, 0, 1, -3}, 0)
}
