package day15

import "testing"

func TestNext(t *testing.T) {
	a, b := examples()
	for i, tt := range []struct {
		expa, expb int64
	}{
		{1092455, 430625591},
		{1181022009, 1233683848},
		{245556042, 1431495498},
		{1744312007, 137874439},
		{1352636452, 285222916},
	} {
		a.next()
		b.next()
		if a.value != tt.expa {
			t.Errorf("i=%d a.next().value => %d, want %d", i, a.value, tt.expa)
		}
		if b.value != tt.expb {
			t.Errorf("i=%d b.next().value => %d, want %d", i, b.value, tt.expb)
		}
	}
}

func TestPickyNext(t *testing.T) {
	a, b := examples()
	a.picky = 4
	b.picky = 8
	for i, tt := range []struct {
		expa, expb int64
	}{
		{1352636452, 1233683848},
		{1992081072, 862516352},
		{530830436, 1159784568},
		{1980017072, 1616057672},
		{740335192, 412269392},
	} {
		a.next()
		b.next()
		if a.value != tt.expa {
			t.Errorf("i=%d a.next().value => %d, want %d", i, a.value, tt.expa)
		}
		if b.value != tt.expb {
			t.Errorf("i=%d b.next().value => %d, want %d", i, b.value, tt.expb)
		}
	}
}

func TestJudge(t *testing.T) {
	for _, tt := range []struct {
		aValue, bValue int64
		iterations     int64
		out            int64
		picky          bool
	}{
		{65, 8921, 5, 1, false},
		//{65, 8921, 40000000, 588, false},
		{65, 8921, 5, 0, true},
		{65, 8921, 1055, 0, true},
		{65, 8921, 1056, 1, true},
		//{65, 8921, 5000000, 309, true},
	} {
		r := Judge(tt.aValue, tt.bValue, tt.iterations, tt.picky)
		if r != tt.out {
			t.Errorf("Judge(%d, %d, %d) => %d, want %d", tt.aValue, tt.bValue, tt.iterations, tt.picky, r, tt.out)
		}
	}
}

func examples() (*generator, *generator) {
	return &generator{16807, 65, 0}, &generator{48271, 8921, 0}
}
