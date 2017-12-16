package day16

import "testing"

func TestMove(t *testing.T) {
	for _, tt := range []struct {
		in   string
		move string
		out  string
	}{
		{"abcde", "s3", "cdeab"},
		{"abcde", "s1", "eabcd"},
		{"eabcd", "x3/4", "eabdc"},
		{"eabdc", "pe/b", "baedc"},
	} {
		r := move(tt.in, tt.move)
		if r != tt.out {
			t.Errorf("Move(%s, %s) => %s, want %s", tt.in, tt.move, r, tt.out)
		}
	}
}

func TestDance(t *testing.T) {
	for _, tt := range []struct {
		in    string
		moves string
		out   string
	}{
		{"abcde", "s1,x3/4,pe/b\n", "baedc"},
	} {
		r := Dance(tt.in, tt.moves)
		if r != tt.out {
			t.Errorf("Dance() => %s, want %s", r, tt.out)
		}
	}
}

func TestLongDance(t *testing.T) {
	for _, tt := range []struct {
		in      string
		moves   string
		repeats int
		out     string
	}{
		{"abcde", "s1,x3/4,pe/b\n", 1, "baedc"},
		{"abcde", "s1,x3/4,pe/b\n", 2, "ceadb"},
		{"abcde", "s1,x3/4,pe/b\n", 19, "ecbda"},
		{"abcde", "s1,x3/4,pe/b\n", 20, "abcde"},
		{"abcde", "s1,x3/4,pe/b\n", 21, "baedc"},
	} {
		r := LongDance(tt.in, tt.moves, tt.repeats)
		if r != tt.out {
			t.Errorf("LongDance(repeats: %d) => %s, want %s", tt.repeats, r, tt.out)
		}
	}
}
