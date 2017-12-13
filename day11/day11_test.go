package day11

import "testing"

func TestMove(t *testing.T) {
	for _, tt := range []struct {
		in  pos
		dir int
		out pos
	}{
		// origin
		{pos{0, 0}, n, pos{1, 0}},
		{pos{0, 0}, ne, pos{1, 1}},
		{pos{0, 0}, se, pos{1, 2}},
		{pos{0, 0}, s, pos{1, 3}},
		{pos{0, 0}, sw, pos{1, 4}},
		{pos{0, 0}, nw, pos{1, 5}},
		// north pole
		{pos{3, 0}, n, pos{4, 0}},
		{pos{3, 0}, ne, pos{4, 1}},
		{pos{3, 0}, se, pos{3, 1}},
		{pos{3, 0}, s, pos{2, 0}},
		{pos{3, 0}, sw, pos{3, 17}},
		{pos{3, 0}, nw, pos{4, 23}},
		// ne side
		{pos{3, 1}, n, pos{4, 1}},
		{pos{3, 1}, ne, pos{4, 2}},
		{pos{3, 1}, se, pos{3, 2}},
		{pos{3, 1}, s, pos{2, 1}},
		{pos{3, 1}, sw, pos{2, 0}},
		{pos{3, 1}, nw, pos{3, 0}},
		// ne corner
		{pos{3, 3}, n, pos{4, 3}},
		{pos{3, 3}, ne, pos{4, 4}},
		{pos{3, 3}, se, pos{4, 5}},
		{pos{3, 3}, s, pos{3, 4}},
		{pos{3, 3}, sw, pos{2, 2}},
		{pos{3, 3}, nw, pos{3, 2}},
		// east side
		{pos{3, 4}, n, pos{3, 3}},
		{pos{3, 4}, ne, pos{4, 5}},
		{pos{3, 4}, se, pos{4, 6}},
		{pos{3, 4}, s, pos{3, 5}},
		{pos{3, 4}, sw, pos{2, 3}},
		{pos{3, 4}, nw, pos{2, 2}},
		// se corner
		{pos{3, 6}, n, pos{3, 5}},
		{pos{3, 6}, ne, pos{4, 7}},
		{pos{3, 6}, se, pos{4, 8}},
		{pos{3, 6}, s, pos{4, 9}},
		{pos{3, 6}, sw, pos{3, 7}},
		{pos{3, 6}, nw, pos{2, 4}},
		// se side
		{pos{3, 7}, n, pos{2, 4}},
		{pos{3, 7}, ne, pos{3, 6}},
		{pos{3, 7}, se, pos{4, 9}},
		{pos{3, 7}, s, pos{4, 10}},
		{pos{3, 7}, sw, pos{3, 8}},
		{pos{3, 7}, nw, pos{2, 5}},
		// south pole
		{pos{3, 9}, n, pos{2, 6}},
		{pos{3, 9}, ne, pos{3, 8}},
		{pos{3, 9}, se, pos{4, 11}},
		{pos{3, 9}, s, pos{4, 12}},
		{pos{3, 9}, sw, pos{4, 13}},
		{pos{3, 9}, nw, pos{3, 10}},
		// sw side
		{pos{3, 11}, n, pos{2, 8}},
		{pos{3, 11}, ne, pos{2, 7}},
		{pos{3, 11}, se, pos{3, 10}},
		{pos{3, 11}, s, pos{4, 14}},
		{pos{3, 11}, sw, pos{4, 15}},
		{pos{3, 11}, nw, pos{3, 12}},
		// sw corner
		{pos{3, 12}, n, pos{3, 13}},
		{pos{3, 12}, ne, pos{2, 8}},
		{pos{3, 12}, se, pos{3, 11}},
		{pos{3, 12}, s, pos{4, 15}},
		{pos{3, 12}, sw, pos{4, 16}},
		{pos{3, 12}, nw, pos{4, 17}},
		// west side
		{pos{3, 14}, n, pos{3, 15}},
		{pos{3, 14}, ne, pos{2, 10}},
		{pos{3, 14}, se, pos{2, 9}},
		{pos{3, 14}, s, pos{3, 13}},
		{pos{3, 14}, sw, pos{4, 18}},
		{pos{3, 14}, nw, pos{4, 19}},
		// nw corner
		{pos{3, 15}, n, pos{4, 21}},
		{pos{3, 15}, ne, pos{3, 16}},
		{pos{3, 15}, se, pos{2, 10}},
		{pos{3, 15}, s, pos{3, 14}},
		{pos{3, 15}, sw, pos{4, 19}},
		{pos{3, 15}, nw, pos{4, 20}},
		// nw side
		{pos{3, 16}, n, pos{4, 22}},
		{pos{3, 16}, ne, pos{3, 17}},
		{pos{3, 16}, se, pos{2, 11}},
		{pos{3, 16}, s, pos{2, 10}},
		{pos{3, 16}, sw, pos{3, 15}},
		{pos{3, 16}, nw, pos{4, 21}},
		// nw side and next to north pole
		{pos{3, 17}, n, pos{4, 23}},
		{pos{3, 17}, ne, pos{3, 0}},
		{pos{3, 17}, se, pos{2, 0}},
		{pos{3, 17}, s, pos{2, 11}},
		{pos{3, 17}, sw, pos{3, 16}},
		{pos{3, 17}, nw, pos{4, 22}},
	} {
		r := tt.in.move(tt.dir)
		if !r.eq(tt.out) {
			t.Errorf("%v.move(%d) => %v, want %v", tt.in, tt.dir, r, tt.out)
		}
	}
}

func TestDistance(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{
		{"ne,ne,ne", 3},
		{"ne,ne,sw,sw", 0},
		{"ne,ne,s,s", 2},
		{"se,sw,se,sw,sw", 3},
	} {
		r := ShortestDistance(tt.in)
		if r != tt.out {
			t.Errorf("ShortestDistance(%s) => %d, want %d", tt.in, r, tt.out)
		}
	}
}
