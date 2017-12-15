package day14

import (
	"fmt"
	"math/bits"
	"strconv"

	"github.com/timocp/adventofcode/day10"
)

func SquaresUsed(input string) int {
	used := 0
	for row := 0; row < 128; row++ {
		for _, b := range day10.HashValue(fmt.Sprintf("%s-%d", input, row)) {
			used += bits.OnesCount8(uint8(b))
		}
	}
	return used
}

type cell struct {
	value  bool // true if corresponding hash bit is 1
	seen   bool // true if we've already considered this cell
	region int  // > 0 if cell is part of a contiguous region
}

const size = 128

type grid struct {
	cells   [size][size]cell
	regions int
}

func (g *grid) value(x, y int) bool {
	return g.cells[x][y].value
}

func (g *grid) seen(x, y int) bool {
	if x < 0 || x >= size || y < 0 || y >= size {
		return true
	}
	return g.cells[x][y].seen
}

func (g *grid) region(x, y int) int {
	return g.cells[x][y].region
}

// marks x,y as current region, and all neighbours recursively
func (g *grid) markRegion(x, y int) {
	g.cells[x][y].seen = true
	g.cells[x][y].region = g.regions

	// in each direction, if not seen and has a value, mark as same region
	if !g.seen(x+1, y) && g.value(x+1, y) {
		g.markRegion(x+1, y)
	}
	if !g.seen(x-1, y) && g.value(x-1, y) {
		g.markRegion(x-1, y)
	}
	if !g.seen(x, y+1) && g.value(x, y+1) {
		g.markRegion(x, y+1)
	}
	if !g.seen(x, y-1) && g.value(x, y-1) {
		g.markRegion(x, y-1)
	}
}

func (g *grid) markSeen(x, y int) {
	g.cells[x][y].seen = true
}

func (g *grid) setTrue(x, y int) {
	g.cells[x][y].value = true
}

func (g *grid) printValues() string {
	s := ""
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if g.value(row, col) {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func (g *grid) printRegions() string {
	s := ""
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if col > 0 {
				s += " "
			}
			if g.value(row, col) {
				s += fmt.Sprintf("%2s", strconv.FormatInt(int64(g.region(row, col)), 36))
			} else {
				s += ".."
			}
		}
		s += "\n"
	}
	return s
}

// RegionsPresent counts the number of distinct regions in the hash based on
// input string
func RegionsPresent(input string) int {
	g := new(grid)

	// mark grids which correspond to bits in the hashes
	for row := 0; row < size; row++ {
		hash := day10.HashValue(fmt.Sprintf("%s-%d", input, row))
		for i, v := range hash {
			for bit := uint(0); bit < 8; bit++ {
				if v&(1<<bit) == (1 << bit) {
					g.setTrue(row, (i+1)*8-int(bit)-1)
				}
			}
		}
	}
	//fmt.Printf("%s", g.printValues())

	// iterate over all cells, marking regions as unfound values found
	for row := 0; row < size; row++ {
		for col := 0; col < size; col++ {
			if !g.seen(row, col) {
				if g.value(row, col) {
					g.regions++
					g.markRegion(row, col)
				} else {
					g.markSeen(row, col)
				}
			}
		}
	}
	//fmt.Printf("%s", g.printRegions())

	return g.regions
}
