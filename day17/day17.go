package day17

import "fmt"

type SpinLock struct {
	buffer []int
	curpos int
	step   int
	next   int
}

func (s *SpinLock) String() string {
	str := ""
	for i, v := range s.buffer {
		if i == s.curpos {
			str += fmt.Sprintf("(%d)", v)
		} else {
			str += fmt.Sprintf(" %d ", v)
		}
	}
	return str
}

func (s *SpinLock) spin() {
	//fmt.Printf("spin(): s=%v\n", s)
	s.curpos = (s.curpos+s.step)%(len(s.buffer)) + 1
	s.buffer = append(s.buffer, 0)
	copy(s.buffer[s.curpos:], s.buffer[s.curpos-1:])
	s.buffer[s.curpos] = s.next
	s.next++
}

func (s *SpinLock) Spins(times int) {
	for i := 0; i < times; i++ {
		s.spin()
	}
}

func (s *SpinLock) AtRel(rel int) int {
	return s.buffer[(s.curpos+rel)%len(s.buffer)]
}

func NewSpinLock(step int) *SpinLock {
	return &SpinLock{[]int{0}, 0, step, 1}
}
