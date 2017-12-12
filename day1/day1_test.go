package day1

import "testing"

func TestInverseCaptcha(t *testing.T) {
	for _, tt := range []struct {
		in   string
		mode int
		out  int
	}{
		{"1122", 0, 3},
		{"1111", 0, 4},
		{"1234", 0, 0},
		{"91212129", 0, 9},
		{"1212", 1, 6},
		{"1221", 1, 0},
		{"123425", 1, 4},
		{"123123", 1, 12},
		{"12131415", 1, 4},
	} {
		r := InverseCaptcha(tt.in, tt.mode)
		if r != tt.out {
			t.Errorf("InverseCaptcha(%s, %d) => %d, want %d", tt.in, tt.mode, r, tt.out)
		}
	}
}
