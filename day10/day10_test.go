package day10

import "testing"

func TestAct(t *testing.T) {
	c := &circle{[]int{0, 1, 2, 3, 4}, 0, 0}
	for i, tt := range []struct {
		length   int
		list     []int
		position int
		skipSize int
	}{
		{3, []int{2, 1, 0, 3, 4}, 3, 1},
		{4, []int{4, 3, 0, 1, 2}, 3, 2},
		{1, []int{4, 3, 0, 1, 2}, 1, 3},
		{5, []int{3, 4, 2, 1, 0}, 4, 4},
	} {
		c.act(tt.length)
		if !eqInts(c.list, tt.list) {
			t.Errorf("%dth act() => %v, want %v", i, c.list, tt.list)
		}
		if c.position != tt.position {
			t.Errorf("%dth act() position => %d, want %d", i, c.position, tt.position)
		}
		if c.skipSize != tt.skipSize {
			t.Errorf("%dth act() skipSize => %d, want %d", i, c.skipSize, tt.skipSize)
		}
	}
}

func TestHash(t *testing.T) {
	r := Hash(5, []int{3, 4, 1, 5})
	if r != 12 {
		t.Errorf("Hash() => %d, want %d", r, 12)
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
