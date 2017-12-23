package day22

import (
	"bufio"
	"io"
)

type grid struct {
	infected   map[[2]int]bool // key is [row, col]
	bursts     int
	infections int
	curpos     [2]int
	direction  int // 0 = up, 1 = right, etc
}

// input is assumed to be an odd width square grid, with center at (0,0)
func loadGrid(input io.Reader) *grid {
	g := &grid{}
	g.infected = make(map[[2]int]bool)
	s := bufio.NewScanner(input)
	offset := 0
	for row := 0; s.Scan(); row++ {
		line := s.Text()
		if row == 0 {
			offset = (len(line) - 1) / 2
		}
		for col, b := range line {
			if b == '#' {
				g.infected[[2]int{row - offset, col - offset}] = true
			}
		}
	}
	if s.Err() != nil {
		panic(s.Err())
	}
	return g
}

func (g *grid) burst() {
	g.bursts++
	if g.infected[g.curpos] {
		g.direction = (g.direction + 1) % 4
	} else {
		g.direction = (g.direction + 3) % 4
	}
	g.infected[g.curpos] = !g.infected[g.curpos]
	if g.infected[g.curpos] {
		g.infections++
	}
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

func Part1(input io.Reader) int {
	g := loadGrid(input)
	for g.bursts < 10000 {
		g.burst()
	}
	return g.infections
}
