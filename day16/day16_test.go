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
