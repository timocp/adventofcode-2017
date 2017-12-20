package day20

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type vector struct {
	x, y, z int
}

type particle struct {
	pos, vel, acc vector
	present       bool
}

func (p *particle) distOrigin() int {
	return abs(p.pos.x) + abs(p.pos.y) + abs(p.pos.z)
}

type Swarm struct {
	p          []*particle
	collisions bool
}

func NewSwarm(in string, collisions bool) (*Swarm, error) {
	s := &Swarm{[]*particle{}, collisions}
	err := s.load(in)
	return s, err
}

func (s *Swarm) Closest() (minP int) {
	// who is closest to origin?
	minDist := math.MaxInt64
	for i := 0; i < len(s.p); i++ {
		if s.p[i].present {
			dist := s.p[i].distOrigin()
			if dist < minDist {
				minDist = dist
				minP = i
			}
		}
	}
	return
}

// Run for n iterations
func (s *Swarm) Run(n int) {
	for i := 0; i < n; i++ {
		s.tick()
	}
}

func (s *Swarm) CountPresent() int {
	count := 0
	for _, p := range s.p {
		if p.present {
			count++
		}
	}
	return count
}

// adjust all velocities and positions
func (s *Swarm) tick() {
	positions := map[vector][]int{}
	for i, p := range s.p {
		if p.present {
			p.vel.x += p.acc.x
			p.vel.y += p.acc.y
			p.vel.z += p.acc.z
			p.pos.x += p.vel.x
			p.pos.y += p.vel.y
			p.pos.z += p.vel.z
			if s.collisions {
				positions[p.pos] = append(positions[p.pos], i)
			}
		}
	}
	if s.collisions {
		// every hash value with more than 1 entry is a set of colliding particles
		for _, v := range positions {
			if len(v) > 1 {
				for _, i := range v {
					s.p[i].present = false
				}
			}
		}
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
			true}
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
