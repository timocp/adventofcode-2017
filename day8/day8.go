package day8

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	target     string
	op         string
	arg        int
	condTarget string
	condOp     string
	condArg    int
}

// Program is a list of instructions
type Program []*instruction

// Registers is a hash of known registers by name with their current value
type Registers map[string]int

// ReadProgram reads the instructions from input and returns a Program
func ReadProgram(input io.Reader) (Program, error) {
	p := []*instruction{}
	s := bufio.NewScanner(input)
	for s.Scan() {
		inst, err := parseInstruction(s.Text())
		if err != nil {
			return p, err
		}
		if !inst.valid() {
			return p, fmt.Errorf("Invalid instruction: %v", inst)
		}
		p = append(p, inst)
	}
	return p, s.Err()
}

// NewRegisters returns a new empty set of registers
func NewRegisters() Registers {
	r := make(map[string]int)
	return r
}

// Execute runs program p against registers r
func (p Program) Execute(r *Registers) {
	for _, inst := range p {
		if inst.test(r) {
			inst.do(r)
		}
	}
}

func (inst *instruction) test(r *Registers) bool {
	condTarget := (*r)[inst.condTarget]
	switch inst.condOp {
	case "<":
		return condTarget < inst.condArg
	case ">":
		return condTarget > inst.condArg
	case "<=":
		return condTarget <= inst.condArg
	case ">=":
		return condTarget >= inst.condArg
	case "==":
		return condTarget == inst.condArg
	case "!=":
		return condTarget != inst.condArg
	}
	panic(fmt.Sprintf("invalid instruction %v", inst))
}

func (inst *instruction) do(r *Registers) {
	switch inst.op {
	case "inc":
		(*r)[inst.target] += inst.arg
	case "dec":
		(*r)[inst.target] -= inst.arg
	default:
		panic(fmt.Sprintf("invalid instruction %v", inst))
	}
}

// LargestValue returns the maximum value found in this set of registers
func (r *Registers) LargestValue() int {
	max := 0
	for _, v := range *r {
		if v > max {
			max = v
		}
	}
	return max
}

var instructionRe = regexp.MustCompile(`^(\w+)\s+(\w+)\s+(-?\d+)\s+if\s+(\w+)\s+(\S+)\s+(-?\d)$`)

func parseInstruction(line string) (*instruction, error) {
	words := strings.Split(line, " ")
	if len(words) != 7 {
		return nil, fmt.Errorf("Malformed instruction (need 7 elements): %s", line)
	}
	if words[3] != "if" {
		return nil, fmt.Errorf("Malformed instruction (missing if): %s", line)
	}
	arg, err := strconv.Atoi(words[2])
	condArg, err2 := strconv.Atoi(words[6])
	if err != nil || err2 != nil {
		return nil, fmt.Errorf("Malformed instruction (op and condOp must an integer): %s", line)
	}
	return &instruction{words[0], words[1], arg, words[4], words[5], condArg}, nil
}

func (inst *instruction) valid() bool {
	if inst.op != "inc" && inst.op != "dec" {
		return false
	}
	if inst.condOp != "<" && inst.condOp != ">" &&
		inst.condOp != "<=" && inst.condOp != ">=" &&
		inst.condOp != "==" && inst.condOp != "!=" {
		return false
	}
	return true
}
