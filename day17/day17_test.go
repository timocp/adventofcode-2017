package day17

import "testing"

func TestSpin(t *testing.T) {
	s := NewSpinLock(3)
	for i, tt := range []string{
		"0 (1)",
		"0 (2) 1",
		"0  2 (3) 1",
		"0  2 (4) 3  1",
		"0 (5) 2  4  3  1",
		"0  5  2  4  3 (6) 1",
		"0  5 (7) 2  4  3  6  1",
		"0  5  7  2  4  3 (8) 6  1",
		"0 (9) 5  7  2  4  3  8  6  1",
	} {
		s.spin()
		if s.String() != tt {
			t.Errorf("%dth spin() => %s, want %s", i, s.String(), tt)
		}
	}
}

func TestSpins(t *testing.T) {
	s := NewSpinLock(3)
	s.Spins(2017)
	if s.AtRel(1) != 638 {
		t.Errorf("2017th spin(), item at curpos+1 = %d, want %d", s.AtRel(1), 638)
	}
}
