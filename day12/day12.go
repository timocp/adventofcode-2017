package day12

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"regexp"
	"strconv"
	"strings"
)

// ProgramList is an array (indexed by program number); each value is the list
// of programs it can talk to
type ProgramList [][]int

var programRe = regexp.MustCompile(`^(\d+)\s*<->\s*([\d,\s]+)$`)

func parseLine(input string) (int, []int, error) {
	if !programRe.MatchString(input) {
		return -1, []int{}, fmt.Errorf("parseLine: Invalid format: %s", input)
	}
	matches := programRe.FindStringSubmatch(input)
	prognum, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1, []int{}, err
	}
	pipes, err := splitInts(matches[2])
	if err != nil {
		return -1, []int{}, err
	}
	return prognum, pipes, nil
}

func splitInts(s string) ([]int, error) {
	words := strings.Split(s, ",")
	ints := make([]int, len(words))
	for i, w := range words {
		tmp, err := strconv.Atoi(strings.TrimSpace(w))
		if err != nil {
			return []int{}, err
		}
		ints[i] = tmp
	}
	return ints, nil
}

// ReadProgramList reads the list of programs and pipes from the input file
func ReadProgramList(input io.Reader) (ProgramList, error) {
	p := ProgramList{}
	s := bufio.NewScanner(input)
	for s.Scan() {
		_, pipes, err := parseLine(s.Text())
		if err != nil {
			return p, err
		}
		p = append(p, pipes)
	}
	if s.Err() != nil {
		return p, s.Err()
	}
	return p, nil
}

// MustReadProgramList calls ReadProgramList and dies if there was an error
func MustReadProgramList(input io.Reader) ProgramList {
	p, err := ReadProgramList(input)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

// iterate through all the pipes we can reach, calling f for each one once
func (p ProgramList) iterate(prognum int, seen map[int]bool, f func(int)) {
	f(prognum)
	seen[prognum] = true
	for _, other := range p[prognum] {
		if !seen[other] {
			p.iterate(other, seen, f)
		}
	}
}

// CountConnected counts the number of programs that can communicate with
// a program by trying all available paths
func (p ProgramList) CountConnected(prognum int) int {
	count := 0
	p.iterate(prognum, make(map[int]bool), func(prognum int) {
		count++
	})
	return count
}

// CountGroups counts the number of groups of programs which can't communicate
// with each other
func (p ProgramList) CountGroups() int {
	seen := make(map[int]bool)
	groups := 0
	for len(seen) < len(p) {
		// first unseen number?
		for i := 0; i < len(p); i++ {
			if !seen[i] {
				groups++
				p.iterate(i, seen, func(prognum int) {})
			}
		}
	}
	return groups
}
