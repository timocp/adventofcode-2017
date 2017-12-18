package day18

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type instruction struct {
	op string // operand
	a1 string // first arg (register or value)
	a2 string // second arg (register or value)
}

type duet struct {
	register  map[string]int
	code      []instruction
	p         int
	lastSound int
}

func newDuet() *duet {
	d := new(duet)
	d.register = make(map[string]int)
	return d
}

func (d *duet) addInstruction(s string) error {
	words := strings.Split(s, " ")
	inst := instruction{words[0], words[1], ""}
	if len(words) == 3 {
		inst.a2 = words[2]
	}
	d.code = append(d.code, inst)
	return nil
}

func (d *duet) set(a string, i int) {
	d.register[a] = i
}

// get value; a can be a register name or a number
func (d *duet) get(a string) int {
	if a[0] >= 'a' && a[0] <= 'z' {
		return d.register[a]
	}
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func (d *duet) load(input io.Reader) error {
	s := bufio.NewScanner(input)
	for s.Scan() {
		d.addInstruction(s.Text())
	}
	return s.Err()
}

// exec runs the next instruction. updates program pointer
// callback frcv is called if a sound is recovered
func (d *duet) exec(frcv func(int)) {
	inst := d.code[d.p]
	fmt.Printf("inst(%d)=%v\treg=%v\n", d.p, inst, d.register)
	switch inst.op {
	case "snd":
		d.lastSound = d.get(inst.a1)
	case "set":
		d.set(inst.a1, d.get(inst.a2))
	case "add":
		d.set(inst.a1, d.get(inst.a1)+d.get(inst.a2))
	case "mul":
		d.set(inst.a1, d.get(inst.a1)*d.get(inst.a2))
	case "mod":
		d.set(inst.a1, d.get(inst.a1)%d.get(inst.a2))
	case "rcv":
		if d.get(inst.a1) != 0 {
			frcv(d.lastSound)
		}
	case "jgz":
		if d.get(inst.a1) > 0 {
			d.p += d.get(inst.a2) - 1
		}
	}
	d.p++
}

// Part1 runs until an rcv is executed.
func Part1(input io.Reader) int {
	d := newDuet()
	err := d.load(input)
	if err != nil {
		log.Fatal(err)
	}
	run := true
	lastSound := 0

	for run {
		d.exec(func(sound int) {
			lastSound = sound
			run = false
		})
	}
	return lastSound
}
