package day9

import (
	"bytes"
	"testing"
)

func TestScore(t *testing.T) {
	for _, tt := range []struct {
		input string
		score int
	}{
		{"{}", 1},
		{"{{{}}}", 6},
		{"{{},{}}", 5},
		{"{{{},{},{{}}}}", 16},
		{"{<a>,<a>,<a>,<a>}", 1},
		{"{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
		{"{<{},{},{{}}>}", 1},
		{"{{},<>}\n", 3},
	} {
		r, err := Score(bytes.NewBufferString(tt.input))
		if err != nil {
			t.Errorf("Score(%s) error %v, want nil", tt.input, err)
		} else if r != tt.score {
			t.Errorf("Score(%s) => %d, want %d", tt.input, r, tt.score)
		}
	}
}
