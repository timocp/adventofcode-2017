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

// first matched transition gets actioned
var transitions = []struct {
	from    int  // current state
	next    byte // input byte
	any     bool // true if any byte is allowed
	delta   int  // how much this action changes the depth
	garbage bool // true if this byte counts as garbage
	to      int  // state to change to
}{
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

type stateMachine struct {
	state   int
	depth   int
	garbage int
}

func (machine *stateMachine) String() string {
	return fmt.Sprintf("[state=%d, depth=%d]", machine.state, machine.depth)
}

func (machine *stateMachine) action(b byte) int {
	for _, transition := range transitions {
		if machine.state == transition.from {
			if transition.any || transition.next == b {
				machine.depth += transition.delta
				if transition.garbage {
					machine.garbage++
				}
				machine.state = transition.to
				return transition.delta
			}
		}
	}
	panic(fmt.Sprintf("no action for machine=%v, b=%q", machine, b))
}

// Result contains statistics after processing
type Result struct {
	Score        int
	GarbageCount int
}

// Process parses the input stream and returns a struct containing stats
func Process(input io.Reader) (Result, error) {
	buf := make([]byte, 1)
	machine := &stateMachine{}
	result := Result{}
	for {
		n, err := input.Read(buf)
		if n > 0 && !unicode.IsSpace(rune(buf[0])) {
			delta := machine.action(buf[0])
			if delta > 0 {
				// if depth is increasing, we've entered a new group.  its score
				// is the current depth
				result.Score += machine.depth
			}
		}
		if err == io.EOF {
			result.GarbageCount = machine.garbage
			return result, nil
		} else if err != nil {
			return result, err
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
