package day17

import "testing"

func TestSpin(t *testing.T) {
	s := NewSpinLock(3)
	for i, tt := range []struct {
		buffer []int
		curpos int
	}{
		{[]int{0, 1}, 1},
		{[]int{0, 2, 1}, 1},
		{[]int{0, 2, 3, 1}, 2},
		{[]int{0, 2, 4, 3, 1}, 2},
		{[]int{0, 5, 2, 4, 3, 1}, 1},
		{[]int{0, 5, 2, 4, 3, 6, 1}, 5},
		{[]int{0, 5, 7, 2, 4, 3, 6, 1}, 2},
		{[]int{0, 5, 7, 2, 4, 3, 8, 6, 1}, 6},
		{[]int{0, 9, 5, 7, 2, 4, 3, 8, 6, 1}, 1},
	} {
		s.spin()
		if !eqInts(s.buffer, tt.buffer) {
			t.Errorf("%dth spin() buffer => %v, want %v", i, s.buffer, tt.buffer)
		}
		if s.curpos != tt.curpos {
			t.Errorf("%dth spin() curpos => %d, want %d", i, s.curpos, tt.curpos)
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

func eqInts(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
