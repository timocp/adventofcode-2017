package day11

import (
	"fmt"
	"os"
	"strings"
)

// position on the hex grid uses a coordinate system (x, y) where
// x is the distance from the origin,
// y is the number of squares around that ring, with x,0 being north and
// counting clockwise.
// there are x*6 squares in each ring (except ring 0 which has 1 square).

const (
	n = iota
	ne
	se
	s
	sw
	nw
)

var dirs = map[string]int{
	"n":  n,
	"ne": ne,
	"se": se,
	"s":  s,
	"sw": sw,
	"nw": nw,
}

type pos struct {
	x int
	y int
}

// return new coordinates when moving dir from current position
// with more time could probably be made more concise, as there is significant
// overlap between the sides.  but at least this is simple to understand.
//
// inverting the conditions (switch on dir at outside) might be better, as
// then the rules can be shared between adjacent sections
func (p pos) move(dir int) pos {
	if p.x == 0 { // origin
		return pos{1, dir}
	} else if p.y == 0 { // north pole
		switch dir {
		case n:
			return pos{p.x + 1, p.y}
		case ne:
			return pos{p.x + 1, p.y + 1}
		case se:
			return pos{p.x, p.y + 1}
		case s:
			return pos{p.x - 1, p.y}
		case sw:
			return pos{p.x, 6*p.x - 1}
		case nw:
			return pos{p.x + 1, 6*(p.x+1) - 1}
		}
	} else if p.y < p.x { // ne side
		switch dir {
		case n:
			return pos{p.x + 1, p.y}
		case ne:
			return pos{p.x + 1, p.y + 1}
		case se:
			return pos{p.x, p.y + 1}
		case s:
			return pos{p.x - 1, p.y}
		case sw:
			return pos{p.x - 1, p.y - 1}
		case nw:
			return pos{p.x, p.y - 1}
		}
	} else if p.y == p.x { // ne corner
		switch dir {
		case n:
			return pos{p.x + 1, p.y}
		case ne:
			return pos{p.x + 1, p.y + 1}
		case se:
			return pos{p.x + 1, p.y + 2}
		case s:
			return pos{p.x, p.y + 1}
		case sw:
			return pos{p.x - 1, p.y - 1}
		case nw:
			return pos{p.x, p.y - 1}
		}
	} else if p.y < p.x*2 { // east side
		switch dir {
		case n:
			return pos{p.x, p.y - 1}
		case ne:
			return pos{p.x + 1, p.y + 1}
		case se:
			return pos{p.x + 1, p.y + 2}
		case s:
			return pos{p.x, p.y + 1}
		case sw:
			return pos{p.x - 1, p.y - 1}
		case nw:
			return pos{p.x - 1, p.y - 2}
		}
	} else if p.y == p.x*2 { // se corner
		switch dir {
		case n:
			return pos{p.x, p.y - 1}
		case ne:
			return pos{p.x + 1, p.y + 1}
		case se:
			return pos{p.x + 1, p.y + 2}
		case s:
			return pos{p.x + 1, p.y + 3}
		case sw:
			return pos{p.x, p.y + 1}
		case nw:
			return pos{p.x - 1, p.y - 2}
		}
	} else if p.y < p.x*3 { // se side
		switch dir {
		case n:
			return pos{p.x - 1, p.y - 3}
		case ne:
			return pos{p.x, p.y - 1}
		case se:
			return pos{p.x + 1, p.y + 2}
		case s:
			return pos{p.x + 1, p.y + 3}
		case sw:
			return pos{p.x, p.y + 1}
		case nw:
			return pos{p.x - 1, p.y - 2}
		}
	} else if p.y == p.x*3 { // south pole
		switch dir {
		case n:
			return pos{p.x - 1, p.y - 3}
		case ne:
			return pos{p.x, p.y - 1}
		case se:
			return pos{p.x + 1, p.y + 2}
		case s:
			return pos{p.x + 1, p.y + 3}
		case sw:
			return pos{p.x + 1, p.y + 4}
		case nw:
			return pos{p.x, p.y + 1}
		}
	} else if p.y < p.x*4 { // sw side
		switch dir {
		case n:
			return pos{p.x - 1, p.y - 3}
		case ne:
			return pos{p.x - 1, p.y - 4}
		case se:
			return pos{p.x, p.y - 1}
		case s:
			return pos{p.x + 1, p.y + 3}
		case sw:
			return pos{p.x + 1, p.y + 4}
		case nw:
			return pos{p.x, p.y + 1}
		}
	} else if p.y == p.x*4 { // sw corner
		switch dir {
		case n:
			return pos{p.x, p.y + 1}
		case ne:
			return pos{p.x - 1, p.y - 4}
		case se:
			return pos{p.x, p.y - 1}
		case s:
			return pos{p.x + 1, p.y + 3}
		case sw:
			return pos{p.x + 1, p.y + 4}
		case nw:
			return pos{p.x + 1, p.y + 5}
		}
	} else if p.y < p.x*5 { // west side
		switch dir {
		case n:
			return pos{p.x, p.y + 1}
		case ne:
			return pos{p.x - 1, p.y - 4}
		case se:
			return pos{p.x - 1, p.y - 5}
		case s:
			return pos{p.x, p.y - 1}
		case sw:
			return pos{p.x + 1, p.y + 4}
		case nw:
			return pos{p.x + 1, p.y + 5}
		}
	} else if p.y == p.x*5 { // nw corner
		switch dir {
		case n:
			return pos{p.x + 1, p.y + 6}
		case ne:
			return pos{p.x, p.y + 1}
		case se:
			return pos{p.x - 1, p.y - 5}
		case s:
			return pos{p.x, p.y - 1}
		case sw:
			return pos{p.x + 1, p.y + 4}
		case nw:
			return pos{p.x + 1, p.y + 5}
		}
	} else { // nw side
		switch dir {
		case n:
			return pos{p.x + 1, p.y + 6}
		case ne:
			return pos{p.x, (p.y + 1) % (p.x * 6)}
		case se:
			return pos{p.x - 1, (p.y - 5) % ((p.x - 1) * 6)}
		case s:
			return pos{p.x - 1, p.y - 6}
		case sw:
			return pos{p.x, p.y - 1}
		case nw:
			return pos{p.x + 1, p.y + 5}
		}
	}
	return p
}

func (p pos) eq(other pos) bool {
	return p.x == other.x && p.y == other.y
}

// ShortstDistance returns the distance to the  origin after following comma
// separate list of moves in path.
func ShortestDistance(path string) int {
	p := pos{0, 0}
	for _, s := range strings.Split(path, ",") {
		if dir, ok := dirs[strings.TrimSpace(s)]; ok {
			p = p.move(dir)
		} else {
			fmt.Fprintf(os.Stderr, "Invalid direction: %s\n", s)
		}
	}
	return p.x
}
