package main

import (
	"fmt"
	"math"
)

// SpiralMemoryDistance measures the manhatten distance from square s to the
// centre of the memory grid
func SpiralMemoryDistance(s int) int {
	r := spiralRadius(s)
	//move := 0
	east := square((r*2)-1) + r

	// work out minimum move in this square to one of the cardinals; can't be
	// more than the spiral radius
	move := r
	for cardinal := 0; cardinal < 4; cardinal++ {
		directPath := east + cardinal*r*2
		thisMove := iAbs(directPath - s)
		if thisMove < move {
			move = thisMove
		}
	}
	return r + move
}

func square(i int) int {
	return i * i
}

func iAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func spiralRadius(s int) int {
	return int((math.Sqrt(float64(s-1)) + 1) / 2)
}

/*
work out which group of squares ("radius" r) away from the middle we're in

r	range
0	1..1
1	2..9
2	10..25
3	26..49
4	50..81
n	(((n*2)-1)^2)+1 .. ((n*2)+1)^2

=> cell x is in square (sqrt(n-1) + 1) / 2 (rounding down)

In square n (n>0), the most direct lines are:

r	east		north		west		south
1	2		4		6		8
2	11		15		19		23
3	28		34		40		46
n	((n*2)-1)**2+n	east+n*2	north+n*2	west+n*2

If on these, the answer is the square number
If not, the minimum difference to one of the cardinals must be moved first
*/

// spiralMemory represents the day 3 memory grid as a slice of slices (2d
// array).  `offset' counts the number of elements which are actually in the
// negative.  The grid is always kept as a square.
type spiralMemory struct {
	offset int
	grid   [][]int
}

// set places the value `i' at coords [x,y].
// if the grid is not big enough, grow() will be called until it is.
func (m *spiralMemory) set(x int, y int, i int) {
	rx := x + m.offset
	ry := y + m.offset
	if ry < 0 || ry > len(m.grid)-1 || rx < 0 || rx > len(m.grid[ry])-1 {
		m.grow()
		m.set(x, y, i)
	} else {
		m.grid[ry][rx] = i
	}
}

// get returns the value currently stored at [x,y].  It returns 0 if the square
// has never been set.
func (m *spiralMemory) get(x int, y int) int {
	rx := x + m.offset
	ry := y + m.offset
	if ry < 0 || ry > len(m.grid)-1 || rx < 0 || rx > len(m.grid[ry])-1 {
		return 0
	}
	return m.grid[ry][rx]
}

// grow causes the grid to become larger.  if the grid is new, it is newly
// allocated as 5x5 ([-2,-2]..[2,2]).
func (m *spiralMemory) grow() {
	if len(m.grid) == 0 {
		// create a new grid (5 squares in each axis)
		m.offset = 2
		m.grid = make([][]int, m.offset*2+1)
		for x := range m.grid {
			m.grid[x] = make([]int, m.offset*2+1)
		}
	} else {
		m.reallocate()
	}
}

// reallocate doubles the size of the grid in both directions; offset is
// doubled.
func (m *spiralMemory) reallocate() {
	// add zero rows before/after the existing ones
	newGrid := make([][]int, m.offset*4+1)
	copy(newGrid[m.offset:], m.grid)
	m.grid = newGrid
	for i := range m.grid {
		if i < 2 || i > m.offset*3 {
			// new blank row
			m.grid[i] = make([]int, m.offset*4+1)
		} else {
			// resize existing row
			newRow := make([]int, m.offset*4+1)
			// copy existing values into the right position, leaving the existing
			// offset's worth of zero values
			copy(newRow[m.offset:], m.grid[i])
			m.grid[i] = newRow
		}
	}
	// record new offset
	m.offset *= 2
}

func (m spiralMemory) String() string {
	s := ""
	for _, row := range m.grid {
		s += "[ "
		for _, square := range row {
			s += fmt.Sprintf("%d ", square)
		}
		s += "]\n"
	}
	return s
}

// sumAdjacent returns the sum of the values in all neighbouring squares of
// [x,y]
func (m *spiralMemory) sumAdjacent(x int, y int) int {
	sum := 0
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			if !(dx == 0 && dy == 0) {
				sum += m.get(x+dx, y+dy)
			}
		}
	}
	return sum
}

// SpiralMemoryStressTest returns the first number larger than i written to
// the memory grid
func SpiralMemoryStressTest(i int) int {
	mem := spiralMemory{}
	mem.set(0, 0, 1)
	x, y := 0, 0
	dir := 0 // 0 is east, 1 is north, etc
	for {
		// move in current dir; turn a turn would be towards an empty square
		switch dir {
		case 0:
			x++
			if mem.get(x, y-1) == 0 {
				dir++
			}
		case 1:
			y--
			if mem.get(x-1, y) == 0 {
				dir++
			}
		case 2:
			x--
			if mem.get(x, y+1) == 0 {
				dir++
			}
		case 3:
			y++
			if mem.get(x+1, y) == 0 {
				dir = 0
			}
		}
		// calculate the new value for this empty square
		v := mem.sumAdjacent(x, y)
		mem.set(x, y, v)
		if v > i {
			return v
		}
	}
}
