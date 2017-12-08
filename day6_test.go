package main

import (
	"bytes"
	"testing"
)

func TestRedistribute(t *testing.T) {
	m := NewMemoryBank([]int{0, 2, 7, 0})
	for i, tt := range []string{
		"[2 4 1 2]",
		"[3 1 2 3]",
		"[0 2 3 4]",
		"[1 3 4 1]",
		"[2 4 1 2]",
	} {
		m.Redistribute()
		if m.String() != tt {
			t.Errorf("After %d redists m => %v, want %s", i, m, tt)
		}
	}
}

func TestRedistributeUntilRepeat(t *testing.T) {
	for _, tt := range []struct {
		in     []int
		state  string
		cycle  int
		repeat int
	}{
		{[]int{0, 2, 7, 0}, "[2 4 1 2]", 5, 4},
	} {
		m := NewMemoryBank(tt.in)
		r := m.RedistributeUntilRepeat()
		if r.Cycle != tt.cycle {
			t.Errorf("RedistributeUntilRepeat().Cycle => %d, want %d", r.Cycle, tt.cycle)
		}
		if r.Repeat != tt.repeat {
			t.Errorf("RedistributeUntilRepeat().Repeat => %d, want %d", r.Repeat, tt.repeat)
		}
		if m.String() != tt.state {
			t.Errorf("Final state => %v, want %s", m, tt.state)
		}
	}
}

func TestNewMemoryBankRead(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out string
	}{
		{"0\t2\t7\t0\n", "[0 2 7 0]"},
	} {
		m, err := NewMemoryBankRead(bytes.NewBufferString(tt.in))
		if err != nil {
			t.Errorf("NewMemoryBankRead(%s) returned error %v, expected nil", tt.in, err)
		} else if m.String() != tt.out {
			t.Errorf("NewMemoryBankRead(%s) => %v, want %s", tt.in, m, tt.out)
		}
	}
}
