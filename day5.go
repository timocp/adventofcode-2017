package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Trampoline represents the instruction set from puzzle 5
type Trampoline struct {
	list []int
	ptr  int
}

// NewTrampoline initialises a trampoline with the given
func NewTrampoline(list []int, ptr int) *Trampoline {
	tmp := make([]int, len(list))
	copy(tmp, list)
	return &Trampoline{list: tmp, ptr: ptr}
}

// NewReadTrampoline creates a new trampoline by reading an instruction per
// line from input
func NewReadTrampoline(input io.Reader) (*Trampoline, error) {
	tramp := &Trampoline{[]int{}, 0}
	s := bufio.NewScanner(input)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return tramp, err
		}
		tramp.list = append(tramp.list, i)
	}
	return tramp, s.Err()
}

// Jump increments the current instruction, then jumps by its original offset
func (tramp *Trampoline) Jump() {
	tramp.list[tramp.ptr]++
	tramp.ptr += tramp.list[tramp.ptr] - 1
}

// StrangeJump is the same but if the current instruction is >=3 it is
// decremented instead of incremented
func (tramp *Trampoline) StrangeJump() {
	offset := tramp.list[tramp.ptr]
	if offset >= 3 {
		tramp.list[tramp.ptr]--
	} else {
		tramp.list[tramp.ptr]++
	}
	tramp.ptr += offset
}

// Escaped returns true if the current pointer is outside the instruction set
func (tramp *Trampoline) Escaped() bool {
	if tramp.ptr < 0 || tramp.ptr >= len(tramp.list) {
		return true
	}
	return false
}

// StepsToExit counts the number of jumps required before the program exits
func (tramp *Trampoline) StepsToExit(strange bool) int {
	count := 0
	for !tramp.Escaped() {
		count++
		if strange {
			tramp.StrangeJump()
		} else {
			tramp.Jump()
		}
	}
	return count
}

func (tramp *Trampoline) String() string {
	s := ""
	for i, instruction := range tramp.list {
		if i == tramp.ptr {
			s += fmt.Sprintf("(%d)", instruction)
		} else {
			s += fmt.Sprintf(" %d ", instruction)
		}
	}
	return s
}
