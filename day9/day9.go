package day9

import (
	"fmt"
	"io"
	"log"
	"unicode"
)

const (
	s0 = iota
	sJunk
	sEscape
)

type transition struct {
	from    int  // current state
	next    byte // input byte
	any     bool // true if any byte is allowed
	delta   int  // how much this action changes the depth
	garbage bool // true if this byte counts as garbage
	to      int  // state to change to
}

// first matched transition gets actioned
var transitions = []transition{
	// s0 = expecting a group but could get junk
	{s0, '{', false, 1, false, s0},
	{s0, '<', false, 0, false, sJunk},
	{s0, ',', false, 0, false, s0},
	{s0, '}', false, -1, false, s0},

	// inside junk.  '>' ends junk, '!' escapes.  all others ignored
	{sJunk, '>', false, 0, false, s0},
	{sJunk, '!', false, 0, false, sEscape},
	{sJunk, 0, true, 0, true, sJunk},

	// inside an escape.  just ignores one character
	{sEscape, 0, true, 0, false, sJunk},
}

// Result contains statistics after processing
type Result struct {
	Score        int
	GarbageCount int
}

type stateMachine struct {
	state  int
	depth  int
	result Result
}

func (machine *stateMachine) String() string {
	return fmt.Sprintf("[state=%d, depth=%d, result=%v]", machine.state, machine.depth, machine.result)
}

func (machine *stateMachine) applyTransition(t transition) {
	machine.state = t.to
	machine.depth += t.delta
	if t.delta > 0 {
		machine.result.Score += machine.depth
	}
	if t.garbage {
		machine.result.GarbageCount++
	}
}

func (machine *stateMachine) action(b byte) {
	for _, t := range transitions {
		if machine.state == t.from {
			if t.any || t.next == b {
				machine.applyTransition(t)
				return
			}
		}
	}
	panic(fmt.Sprintf("no action for machine=%v, b=%q", machine, b))
}

// Process parses the input stream and returns a struct containing stats
func Process(input io.Reader) (Result, error) {
	buf := make([]byte, 1)
	machine := &stateMachine{}
	for {
		n, err := input.Read(buf)
		if n > 0 && !unicode.IsSpace(rune(buf[0])) {
			machine.action(buf[0])
		}
		if err == io.EOF {
			return machine.result, nil
		} else if err != nil {
			return machine.result, err
		}
	}
}

// MustProcess calls Process and dies if there was an error
func MustProcess(input io.Reader) Result {
	r, err := Process(input)
	if err != nil {
		log.Fatal(err)
	}
	return r
}
