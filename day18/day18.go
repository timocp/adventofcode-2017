package day18

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"strconv"
	"strings"
	"sync"
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
	fixed     bool // true if this is a "fixed" program from part 2
	inPipe    chan int
	outPipe   chan int
	sent      int
	waiting   bool
	mutex     sync.Mutex
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

// true if program is runnable (pointer is in bounds)
func (d *duet) runable() bool {
	return d.p >= 0 && d.p < len(d.code)
}

// exec runs the next instruction. updates program pointer
// callback frcv is called if a sound is recovered
func (d *duet) exec(frcv func(int)) {
	inst := d.code[d.p]
	switch inst.op {
	case "snd":
		if d.fixed {
			d.outPipe <- d.get(inst.a1)
			d.sent++
		} else {
			d.lastSound = d.get(inst.a1)
		}
	case "set":
		d.set(inst.a1, d.get(inst.a2))
	case "add":
		d.set(inst.a1, d.get(inst.a1)+d.get(inst.a2))
	case "mul":
		d.set(inst.a1, d.get(inst.a1)*d.get(inst.a2))
	case "mod":
		d.set(inst.a1, d.get(inst.a1)%d.get(inst.a2))
	case "rcv":
		if d.fixed {
			select {
			case tmp := <-d.inPipe:
				d.set(inst.a1, tmp)
				d.waiting = false
			default:
				d.waiting = true
				d.p--
			}
		} else if d.get(inst.a1) != 0 {
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

	for run && d.runable() {
		d.exec(func(sound int) {
			lastSound = sound
			run = false
		})
	}
	return lastSound
}

// makes a new duet and fills in the parts used by part 2
func newFixedDuet(nbr int, input string) *duet {
	p := newDuet()
	p.load(bytes.NewBufferString(input))
	p.fixed = true
	p.outPipe = make(chan int, 100)
	p.set("p", nbr)
	return p
}

func (d *duet) isWaiting() bool {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.waiting
}

func runDuet(p, other *duet, done chan bool) {
	for p.runable() {
		p.exec(func(int) {})
		if p.waiting && other.isWaiting() {
			break
		}
	}
	done <- true
}

func Part2(input string) int {
	p0, p1 := newFixedDuet(0, input), newFixedDuet(1, input)
	p0.inPipe = p1.outPipe
	p1.inPipe = p0.outPipe
	done := make(chan bool, 2)
	go runDuet(p0, p1, done)
	go runDuet(p1, p0, done)
	<-done
	<-done
	return p1.sent
}
