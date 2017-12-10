package day8

import (
	"bytes"
	"testing"
)

var sampleProgram = `b inc 5 if a > 1
a inc 1 if b < 5
c dec -10 if a >= 1
c inc -20 if c == 10`

func TestReadProgram(t *testing.T) {
	p, err := ReadProgram(bytes.NewBufferString(sampleProgram))
	if err != nil {
		t.Errorf("ReadProgram() returned error %v, want nil", err)
	} else if len(p) != 4 {
		t.Errorf("ReadProgram() len => %d, want %d", len(p), 4)
	} else {
		for i, tt := range []*instruction{
			&instruction{"b", "inc", 5, "a", ">", 1},
			&instruction{"a", "inc", 1, "b", "<", 5},
			&instruction{"c", "dec", -10, "a", ">=", 1},
			&instruction{"c", "inc", -20, "c", "==", 10},
		} {
			if !equalInstruction(p[i], tt) {
				t.Errorf("p[%d] => %v, want %v", i, p[i], tt)
			}
		}
	}
}

func TestExecute(t *testing.T) {
	p, _ := ReadProgram(bytes.NewBufferString(sampleProgram))
	r := Registers{}
	max := p.Execute(&r)
	if len(r) != 2 {
		t.Errorf("%d registers after execution, want %d", len(r), 2)
	}
	if max != 10 {
		t.Errorf("Maximum value => %d, want %d", max, 10)
	}
	for _, tt := range []struct {
		name  string
		value int
	}{
		{"a", 1},
		{"b", 0},
		{"c", -10},
	} {
		if r[tt.name] != tt.value {
			t.Errorf("Register %s => %d, want %d", tt.name, r[tt.name], tt.value)
		}
	}
}

func TestLargestValue(t *testing.T) {
	p, _ := ReadProgram(bytes.NewBufferString(sampleProgram))
	r := Registers{}
	p.Execute(&r)
	if r.LargestValue() != 1 {
		t.Errorf("LargestValue() => %d, want %d", r.LargestValue(), 1)
	}
}

func equalInstruction(a, b *instruction) bool {
	if a.target != b.target || a.op != b.op ||
		a.arg != b.arg || a.condTarget != b.condTarget ||
		a.condOp != b.condOp || a.condArg != b.condArg {
		return false
	}
	return true
}
