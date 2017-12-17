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
