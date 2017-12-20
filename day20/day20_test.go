package day20

import "testing"

var in = `p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
`

func TestClosest(t *testing.T) {
	s, err := NewSwarm(in)
	if err != nil {
		t.Errorf("NewSwarm(...) => error %v", err)
	}
	exp := 0
	r := s.Closest()
	if r != exp {
		t.Errorf("Closest() => %d, want %d", r, exp)
	}
}
