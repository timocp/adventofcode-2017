package day3

import "testing"

func TestSpiralMemoryDistance(t *testing.T) {
	for _, tt := range []struct {
		in  int
		out int
	}{
		{1, 0},
		{12, 3},
		{23, 2},
		{1024, 31},
	} {
		r := SpiralMemoryDistance(tt.in)
		if r != tt.out {
			t.Errorf("SpiralMemoryDistance(%d) => %d, want %d", tt.in, r, tt.out)
		}
	}
}

func TestSpiralMemoryStressTest(t *testing.T) {
	for _, tt := range []struct {
		in  int
		out int
	}{
		{1, 2},
		{25, 26},
		{100, 122},
		{800, 806},
	} {
		r := SpiralMemoryStressTest(tt.in)
		if r != tt.out {
			t.Errorf("SpiralMemoryStressTest(%d) => %d, want %d", tt.in, r, tt.out)
		}
	}
}
