package day22

import (
	"bufio"
	"io"
)

const (
	clean = iota
	weakened
	infected
	flagged
)

type grid struct {
	state      map[[2]int]int // key is [row, col]
	bursts     int
	infections int
	curpos     [2]int
	direction  int // 0 = up, 1 = right, etc
}

// input is assumed to be an odd width square grid, with center at (0,0)
func loadGrid(input io.Reader) *grid {
	g := &grid{}
	g.state = make(map[[2]int]int)
	s := bufio.NewScanner(input)
	offset := 0
	for row := 0; s.Scan(); row++ {
		line := s.Text()
		if row == 0 {
			offset = (len(line) - 1) / 2
		}
		for col, b := range line {
			if b == '#' {
				g.state[[2]int{row - offset, col - offset}] = infected
			}
		}
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return g
}

func (g *grid) move() {
	switch g.direction {
	case 0:
		g.curpos[0]--
	case 1:
		g.curpos[1]++
	case 2:
		g.curpos[0]++
	case 3:
		g.curpos[1]--
	}
}

func (g *grid) burst() {
	g.bursts++
	if g.state[g.curpos] == infected {
		g.direction = (g.direction + 1) % 4
		g.state[g.curpos] = clean
	} else {
		g.direction = (g.direction + 3) % 4
		g.state[g.curpos] = infected
		g.infections++
	}
	g.move()
}

func (g *grid) evolvedBurst() {
	g.bursts++
	switch g.state[g.curpos] {
	case clean:
		g.direction = (g.direction + 3) % 4
	case weakened:
		g.infections++
	case infected:
		g.direction = (g.direction + 1) % 4
	case flagged:
		g.direction = (g.direction + 2) % 4
	}
	g.state[g.curpos] = (g.state[g.curpos] + 1) % 4
	g.move()
}

func Run(input io.Reader, n int, evolved bool) int {
	g := loadGrid(input)
	for g.bursts < n {
		if evolved {
			g.evolvedBurst()
		} else {
			g.burst()
		}
	}
	return g.infections
}
