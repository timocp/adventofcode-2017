package day9

import (
	"fmt"
	"io"
	"unicode"
)

const (
	s0 = iota
	sJunk
	sEscape
)

// first matched transition gets actioned
var transitions = []struct {
	from  int  // current state
	next  byte // input byte
	any   bool // true if any byte is allowed
	delta int  // how much this action changes the depth
	to    int  // state to change to
}{
	// s0 = expecting a group but could get junk
	{s0, '{', false, 1, s0},
	{s0, '<', false, 0, sJunk},
	{s0, ',', false, 0, s0},
	{s0, '}', false, -1, s0},

	// inside junk.  '>' ends junk, '!' escapes.  all others ignored
	{sJunk, '>', false, 0, s0},
	{sJunk, '!', false, 0, sEscape},
	{sJunk, 0, true, 0, sJunk},

	// inside an escape.  just ignores one character
	{sEscape, 0, true, 0, sJunk},
}

type stateMachine struct {
	state int
	depth int
}

func (machine *stateMachine) String() string {
	return fmt.Sprintf("[state=%d, depth=%d]", machine.state, machine.depth)
}

func (machine *stateMachine) action(b byte) int {
	for _, transition := range transitions {
		if machine.state == transition.from {
			if transition.any || transition.next == b {
				machine.depth += transition.delta
				machine.state = transition.to
				return transition.delta
			}
		}
	}
	panic(fmt.Sprintf("no action for machine=%v, b=%q", machine, b))
}

// Score parses the input stream and returns its total score, which is the sum
// of all group scores
func Score(input io.Reader) (int, error) {
	buf := make([]byte, 1)
	machine := &stateMachine{}
	score := 0
	for {
		n, err := input.Read(buf)
		if n > 0 && !unicode.IsSpace(rune(buf[0])) {
			delta := machine.action(buf[0])
			if delta > 0 {
				// if depth is increasing, we've entered a new group.  its score
				// is the current depth
				score += machine.depth
			}
		}
		if err == io.EOF {
			return score, nil
		} else if err != nil {
			return score, err
		}
	}
}
