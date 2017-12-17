package day17

import (
	"fmt"
	"strings"
)

type item struct {
	value int
	next  *item
}

type SpinLock struct {
	head    *item
	current *item
	step    int
	next    int
}

func (s *SpinLock) String() string {
	str := ""
	for it := s.head; it != nil; it = it.next {
		if it == s.current {
			str += fmt.Sprintf("(%d)", it.value)
		} else {
			str += fmt.Sprintf(" %d ", it.value)
		}
	}
	return strings.TrimSpace(str)
}

func (s *SpinLock) spin() {
	ptr := s.walk(s.current, s.step)
	s.current = &item{s.next, ptr.next}
	ptr.next = s.current
	s.next++
}

func (s *SpinLock) Spins(times int) {
	for i := 0; i < times; i++ {
		s.spin()
	}
}

// returns item after walking from ptr n times
func (s *SpinLock) walk(ptr *item, n int) *item {
	for i := 0; i < n; i++ {
		if ptr.next == nil {
			ptr = s.head
		} else {
			ptr = ptr.next
		}
	}
	return ptr
}

func (s *SpinLock) At(pos int) int {
	return s.walk(s.head, pos).value
}

func (s *SpinLock) AtRel(rel int) int {
	return s.walk(s.current, rel).value
}

func NewSpinLock(step int) *SpinLock {
	first := &item{0, nil}
	return &SpinLock{first, first, step, 1}
}

// StopValue returns the value after a zero (always 2nd pos) after n spins.
// This is the same as s.Spins(n); s.At(1) without actually performing the skips
func (s *SpinLock) StopValue(n int) int {
	value := 0
	curpos := 0
	for i := 1; i <= n; i++ {
		curpos = ((curpos + s.step) % i) + 1
		if curpos == 1 {
			value = i
		}
	}
	return value
}
