package day20

import "testing"

var in = `p=<3,0,0>, v=<2,0,0>, a=<-1,0,0>
p=<4,0,0>, v=<0,0,0>, a=<-2,0,0>
`

var in2 = `p=<-6,0,0>, v=<3,0,0>, a=<0,0,0>
p=<-4,0,0>, v=<2,0,0>, a=<0,0,0>
p=<-2,0,0>, v=<1,0,0>, a=<0,0,0>
p=<3,0,0>, v=<-1,0,0>, a=<0,0,0>
`

func TestSwarm(t *testing.T) {
	for i, tt := range []struct {
		input   string
		iters   int
		closest int
		present int
	}{
		{in, 3, 0, 2},
		{in2, 3, 3, 1},
	} {
		s, err := NewSwarm(tt.input, true)
		if err != nil {
			t.Fatalf("NewSwarm %d => error %v", i, err)
		}
		s.Run(tt.iters)
		r := s.Closest()
		if r != tt.closest {
			t.Errorf("Swarm %d, Closest() => %d, want %d", i, r, tt.closest)
		}
		r = s.CountPresent()
		if r != tt.present {
			t.Errorf("Swarm %d, CountPresent() => %d, want %d", i, r, tt.present)
		}
	}
}
