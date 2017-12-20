package day20

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type vector struct {
	x, y, z int
}

type particle struct {
	pos, vel, acc vector
}

func (p *particle) distOrigin() int {
	return abs(p.pos.x) + abs(p.pos.y) + abs(p.pos.z)
}

type Swarm struct {
	p []*particle
}

func NewSwarm(in string) (*Swarm, error) {
	s := &Swarm{[]*particle{}}
	err := s.load(in)
	return s, err
}

func (s *Swarm) Closest() int {
	// not sure yet how to detect unchanging "long term" winner. let's just
	// run for a sufficiently long time :)
	for i := 0; i < 1000; i++ {
		s.tick()
	}
	// who is closest to origin?
	minDist := s.p[0].distOrigin()
	minP := 0
	for i := 1; i < len(s.p); i++ {
		dist := s.p[i].distOrigin()
		if dist < minDist {
			minDist = dist
			minP = i
		}
	}
	return minP
}

// adjust all velocities and positions
func (s *Swarm) tick() {
	for _, p := range s.p {
		p.vel.x += p.acc.x
		p.vel.y += p.acc.y
		p.vel.z += p.acc.z
		p.pos.x += p.vel.x
		p.pos.y += p.vel.y
		p.pos.z += p.vel.z
	}
}

var lineRe = regexp.MustCompile(`^p=<(-?\d+),(-?\d+),(-?\d+)>, v=<(-?\d+),(-?\d+),(-?\d+)>, a=<(-?\d+),(-?\d+),(-?\d+)>$`)

func (s *Swarm) load(in string) error {
	var err error
	toInt := func(s string) (i int) {
		if err != nil {
			return
		}
		i, err = strconv.Atoi(s)
		return
	}
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			continue
		}
		if !lineRe.MatchString(line) {
			return fmt.Errorf("Invalid input: %s", line)
		}
		matches := lineRe.FindStringSubmatch(line)
		p := particle{
			vector{toInt(matches[1]), toInt(matches[2]), toInt(matches[3])},
			vector{toInt(matches[4]), toInt(matches[5]), toInt(matches[6])},
			vector{toInt(matches[7]), toInt(matches[8]), toInt(matches[9])},
		}
		s.p = append(s.p, &p)
	}
	return err
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
