package main

import (
	"bytes"
	"testing"
)

func TestCorruptionChecksum(t *testing.T) {
	tt := [][]int{
		{5, 1, 9, 5},
		{7, 5, 3},
		{2, 4, 6, 8},
	}
	r := CorruptionChecksum(tt)
	if r != 18 {
		t.Errorf("CorruptionChecksum(%v) => %d, want %d", tt, r, 18)
	}
}

func TestReadSpreadsheet(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out [][]int
	}{
		{"5 1 9 5\n7 5 3\n2 4 6 8\n", [][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}},
	} {
		r, err := ReadSpreadsheet(bytes.NewBufferString(tt.in))
		if err != nil {
			t.Errorf("ReadSpreadsheet returned error %v, expected nil", err)
		}
		if !compareSS(r, tt.out) {
			t.Errorf("ReadSpreadsheet(%v) => %v, want %v", tt.in, r, tt.out)
		}
	}
}

// returns true iff both 2d slices are the same
func compareSS(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, row := range a {
		if len(row) != len(b[i]) {
			return false
		}
		for j, cell := range row {
			if cell != b[i][j] {
				return false
			}
		}
	}
	return true
}
