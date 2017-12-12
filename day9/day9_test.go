package day9

import (
	"bytes"
	"testing"
)

func TestResult(t *testing.T) {
	for _, tt := range []struct {
		input   string
		score   int
		garbage int
	}{
		{"{}", 1, 0},
		{"{{{}}}", 6, 0},
		{"{{},{}}", 5, 0},
		{"{{{},{},{{}}}}", 16, 0},
		{"{<a>,<a>,<a>,<a>}", 1, 4},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9, 8},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9, 0},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3, 17},
		{"{<{},{},{{}}>}", 1, 10},
		{"{{},<>}\n", 3, 0},
	} {
		r, err := Process(bytes.NewBufferString(tt.input))
		if err != nil {
			t.Errorf("Process(%s) error %v, want nil", tt.input, err)
		} else {
			if r.Score != tt.score {
				t.Errorf("Score(%s) => %d, want %d", tt.input, r.Score, tt.score)
			}
			if r.GarbageCount != tt.garbage {
				t.Errorf("GarbageCount(%s) => %d, want %d", tt.input, r.GarbageCount, tt.garbage)
			}
		}
	}
}
