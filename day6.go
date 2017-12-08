package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// MemoryBank represents the memory bank in puzzle 6
type MemoryBank struct {
	blocks []int
}

// NewMemoryBank creates a MemoryBank initialised with the given integer slice
func NewMemoryBank(blocks []int) *MemoryBank {
	tmp := make([]int, len(blocks))
	copy(tmp, blocks)
	return &MemoryBank{blocks: tmp}
}

// NewMemoryBankRead creates a MemoryBank initialised with integers read from
// input
func NewMemoryBankRead(input io.Reader) (*MemoryBank, error) {
	m := &MemoryBank{[]int{}}
	s := bufio.NewScanner(input)
	for s.Scan() {
		for _, v := range strings.Fields(s.Text()) {
			i, err := strconv.Atoi(v)
			if err != nil {
				return m, err
			}
			m.blocks = append(m.blocks, i)
		}
	}
	return m, s.Err()

}

func (m *MemoryBank) String() string {
	s := "["
	for i, b := range m.blocks {
		if i > 0 {
			s += " "
		}
		s += strconv.Itoa(b)
	}
	return s + "]"
}

// Redistribute picks the memory bank with the most blocks and redistributes
// them to other banks to the right one-by-one
func (m *MemoryBank) Redistribute() {
	// choose index to redistribute
	target := 0
	for i, b := range m.blocks {
		if b > m.blocks[target] {
			target = i
		}
	}

	// Redistribute its blocks
	blocks := m.blocks[target]
	m.blocks[target] = 0
	for blocks > 0 {
		target++
		if target >= len(m.blocks) {
			target = 0
		}
		m.blocks[target]++
		blocks--
	}
}

// RedistributeUntilRepeat calls redistribute until the configuration has been
// seen before.  It returns the number of steps that were taken to do so.
func (m *MemoryBank) RedistributeUntilRepeat() int {
	// keep a hash of the string representations we've seen
	seen := make(map[string]bool)
	count := 0
	for {
		m.Redistribute()
		count++
		state := m.String()
		if seen[state] {
			return count
		}
		seen[state] = true
	}
}
