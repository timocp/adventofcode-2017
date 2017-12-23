package day21

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type grid struct {
	p [][]bool
}

func (g *grid) size() int {
	return len(g.p)
}

func (g *grid) String() string {
	s := ""
	for i, row := range g.p {
		if i > 0 {
			s += "/"
		}
		for _, cell := range row {
			if cell {
				s += "#"
			} else {
				s += "."
			}
		}
	}
	return s
}

func (g *grid) copy() *grid {
	size := g.size()
	q := make([][]bool, size)
	for i := 0; i < size; i++ {
		q[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			q[i][j] = g.p[i][j]
		}
	}
	return &grid{q}
}

// rotate clockwise
func (g *grid) rotate() *grid {
	size := g.size()
	q := make([][]bool, size)
	for i := 0; i < size; i++ {
		q[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			q[i][j] = g.p[size-j-1][i]
		}
	}
	return &grid{q}
}

// flip horizontally
func (g *grid) hflip() *grid {
	size := g.size()
	q := make([][]bool, size)
	for i := 0; i < size; i++ {
		q[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			q[i][j] = g.p[i][size-j-1]
		}
	}
	return &grid{q}
}

// flip vertically
func (g *grid) vflip() *grid {
	size := g.size()
	q := make([][]bool, size)
	for i := 0; i < size; i++ {
		q[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			q[i][j] = g.p[size-i-1][j]
		}
	}
	return &grid{q}
}

func (g *grid) equal(other *grid) bool {
	if g.size() != other.size() {
		return false
	}
	for i, row := range g.p {
		for j, cell := range row {
			if cell != other.p[i][j] {
				return false
			}
		}
	}
	return true
}

// call f for each possible move from here (rotates and flips)
func (g *grid) permutate(f func(*grid)) {
	z := g.copy()
	f(z)
	for i := 0; i < 3; i++ {
		z = z.rotate()
		f(z)
	}
	z = g.hflip()
	f(z)
	for i := 0; i < 3; i++ {
		z = z.rotate()
		f(z)
	}
	z = g.vflip()
	f(z)
	for i := 0; i < 3; i++ {
		z = z.rotate()
		f(z)
	}
}

type rule struct {
	in  *grid
	out *grid
}

func (r rule) String() string {
	return fmt.Sprintf("%v => %v", r.in, r.out)
}

// will be 2x2: ../..
// or 3x3: .../.../...
// or 4x4 (output rules)
func parseGrid(in string) *grid {
	size := 2
	if len(in) == 11 {
		size = 3
	} else if len(in) == 19 {
		size = 4
	}
	p := make([][]bool, size)
	for i, pattern := range strings.Split(in, "/") {
		p[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			if pattern[j] == '#' {
				p[i][j] = true
			}
		}
	}
	return &grid{p}
}

func loadRules(input io.Reader) []rule {
	r := make([]rule, 0)
	s := bufio.NewScanner(input)
	for s.Scan() {
		tmp := strings.Split(s.Text(), " => ")
		in := parseGrid(tmp[0])
		out := parseGrid(tmp[1])
		in.permutate(func(g *grid) {
			r = append(r, rule{g, out})
		})
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return r
}

// extract a subgrid (offset x,y relative to size)
func (g *grid) sub(x, y, size int) *grid {
	q := make([][]bool, size)
	for i := 0; i < size; i++ {
		q[i] = make([]bool, size)
		for j := 0; j < size; j++ {
			q[i][j] = g.p[x+i][y+j]
		}
	}
	return &grid{q}
}

func (g *grid) enhance(rules []rule) *grid {
	for _, r := range rules {
		if g.equal(r.in) {
			return r.out
		}
	}
	panic(fmt.Errorf("No rule for %s", g))
}

func (g *grid) iterate(rules []rule) *grid {
	size := g.size()
	var chunkSize, newChunkSize int
	if size%2 == 0 {
		chunkSize = 2
		newChunkSize = 3
	} else {
		chunkSize = 3
		newChunkSize = 4
	}
	chunks := size / chunkSize
	newSize := chunks * newChunkSize

	// build up a 2d array of the new grids
	tmp := make([][]*grid, chunks)
	for x := 0; x < chunks; x++ {
		tmp[x] = make([]*grid, chunks)
		for y := 0; y < chunks; y++ {
			tmp[x][y] = g.sub(x*chunkSize, y*chunkSize, chunkSize).enhance(rules)
		}
	}

	/*
		fmt.Printf("size=%d, chunkSize=%d / chunks=%d / newSize=%d, newChunkSize=%d\n", size, chunkSize, chunks, newSize, newChunkSize)
		fmt.Printf("before: %v\n", g)
		fmt.Printf("after:  %v\n", tmp)
	*/

	// create target grid
	q := make([][]bool, newSize)
	for i := 0; i < newSize; i++ {
		q[i] = make([]bool, newSize)
	}

	// copy all temporary grids into the right squares
	for x := 0; x < chunks; x++ {
		for y := 0; y < chunks; y++ {
			for i := 0; i < newChunkSize; i++ {
				for j := 0; j < newChunkSize; j++ {
					q[x*newChunkSize+i][y*newChunkSize+j] = tmp[x][y].p[i][j]
				}
			}
		}
	}
	return &grid{q}
}

func (g *grid) pixels() int {
	size := g.size()
	count := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if g.p[i][j] {
				count++
			}
		}
	}
	return count
}

func startingGrid() *grid {
	return parseGrid(".#./..#/###")
}

func IterateAndCountPixels(input io.Reader, n int) int {
	rules := loadRules(input)
	g := startingGrid()
	for i := 0; i < n; i++ {
		g = g.iterate(rules)
	}
	return g.pixels()
}
