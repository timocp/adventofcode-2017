package day19

import (
	"fmt"
	"strings"
)

const (
	north = iota
	east
	south
	west
)

var rowOffset = []int{-1, 0, 1, 0}
var colOffset = []int{0, 1, 0, -1}

type Tubes struct {
	diagram  []string
	row, col int
	dir      int
	letters  string
	stopped  bool
}

func (t *Tubes) String() string {
	s := fmt.Sprintf("-- dir=%d letters=%s -----\n", t.dir, t.letters)
	for i, line := range t.diagram {
		if i == t.row {
			s += line[:t.col]
			s += "*"
			s += line[t.col+1:]
		} else {
			s += line
		}
		s += "\n"
	}
	return s
}

func NewTubes(s string) *Tubes {
	t := &Tubes{}
	t.diagram = strings.Split(s, "\n")
	t.col = strings.Index(t.diagram[0], "|")
	t.dir = south
	return t
}

func (t *Tubes) at(row, col int) byte {
	if row < 0 || row >= len(t.diagram) || col < 0 || col >= len(t.diagram[row]) {
		return byte(' ')
	}
	return t.diagram[row][col]
}

func (t *Tubes) lookAhead() byte {
	//fmt.Printf("lookAhead()\n")
	return t.at(t.row+rowOffset[t.dir], t.col+colOffset[t.dir])
}

func (t *Tubes) lookLeft() byte {
	//fmt.Printf("lookLeft() %d %d\n", (t.dir+3)%4, (t.dir+3)%4)
	return t.at(t.row+rowOffset[(t.dir+3)%4], t.col+colOffset[(t.dir+3)%4])
}

func (t *Tubes) lookRight() byte {
	//fmt.Printf("lookRight()\n")
	return t.at(t.row+rowOffset[(t.dir+1)%4], t.col+colOffset[(t.dir+1)%4])
}

func (t *Tubes) step() {
	// turn right or left only if there is no option
	//fmt.Printf("%v\n", t)
	if t.lookAhead() == byte(' ') {
		if t.lookLeft() != byte(' ') {
			t.dir = (t.dir + 3) % 4
		} else if t.lookRight() != byte(' ') {
			t.dir = (t.dir + 1) % 4
		} else {
			t.stopped = true
			return
		}
	}
	t.row += rowOffset[t.dir]
	t.col += colOffset[t.dir]
	this := t.at(t.row, t.col)
	if this >= 'A' && this <= 'Z' {
		t.letters += string(this)
	}
}

func (t *Tubes) Walk() string {
	for !t.stopped {
		t.step()
	}
	return t.letters
}
