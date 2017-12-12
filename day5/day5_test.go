package day5

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

func TestStrangeJump(t *testing.T) {
	tramp := sample()
	for i, exp := range []string{
		"(1) 3  0  1  -3 ",
		" 2 (3) 0  1  -3 ",
		" 2  2  0  1 (-3)",
		" 2 (2) 0  1  -2 ",
		" 2  3  0 (1) -2 ",
		" 2  3  0  2 (-2)",
		" 2  3 (0) 2  -1 ",
		" 2  3 (1) 2  -1 ",
		" 2  3  2 (2) -1 ",
		" 2  3  2  3  -1 ",
	} {
		if tramp.Escaped() {
			t.Errorf("%dth tramp.Escaped() => true, want false", i+1)
		}
		tramp.StrangeJump()
		if tramp.String() != exp {
			t.Errorf("%dth tramp.StrangeJump() => %s, want %s", i+1, tramp.String(), exp)
			break
		}
	}
	// last jump should escape
	if !tramp.Escaped() {
		t.Errorf("Final tramp.Escaped() => false, want true")
	}
}

func TestStepsToExit(t *testing.T) {
	for _, tt := range []struct {
		strange bool
		out     int
	}{
		{false, 5},
		{true, 10},
	} {
		r := sample().StepsToExit(tt.strange)
		if r != tt.out {
			t.Errorf("%v.StepsToExit(%t) => %d, want %d", sample(), tt.strange, r, tt.out)
		}
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
