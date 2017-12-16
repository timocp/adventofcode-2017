package day13

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type layer struct {
	max int // range of scanner
	pos int
	dir int // 1 down, -1 up
}

type Firewall struct {
	layers   []*layer
	Severity int
	caught   bool
}

func (f *Firewall) String() string {
	s := ""
	for i, layer := range f.layers {
		s += fmt.Sprintf("%2d ", i)
		if layer != nil {
			for j := 0; j < layer.max; j++ {
				if j == layer.pos {
					if layer.dir == 1 {
						s += "[v]"
					} else {
						s += "[^]"
					}
				} else {
					s += "[ ]"
				}
			}
		} else {
			s += "..."
		}
		s += "\n"
	}
	return s
}

func ReadFirewall(input io.Reader) (*Firewall, error) {
	f := &Firewall{}
	s := bufio.NewScanner(input)
	for s.Scan() {
		words := strings.Split(s.Text(), ": ")
		if len(words) != 2 {
			return nil, fmt.Errorf("Invalid format: %s", s.Text())
		}
		depth, err := strconv.Atoi(words[0])
		if err != nil {
			return nil, fmt.Errorf("Invalid depth: %s", s.Text())
		}
		_range, err := strconv.Atoi(words[1])
		if err != nil {
			return nil, fmt.Errorf("Invalid range: %s", s.Text())
		}
		for len(f.layers) < depth {
			// pad empty layers
			f.layers = append(f.layers, nil)
		}
		f.layers = append(f.layers, &layer{_range, 0, 1})
	}
	if s.Err() != nil {
		return nil, s.Err()
	}

	return f, nil
}

// moves all scanners
func (f *Firewall) tick() {
	for _, layer := range f.layers {
		if layer != nil {
			if layer.max == 0 {
				// can't actually move
				layer.dir = -layer.dir
			} else {
				layer.pos += layer.dir
				if layer.pos < 0 {
					layer.pos = 1
					layer.dir = -layer.dir
				} else if layer.pos == layer.max {
					layer.pos = layer.max - 2
					layer.dir = -layer.dir
				}
			}
		}
	}
}

func MustReadFirewall(input io.Reader) *Firewall {
	f, err := ReadFirewall(input)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func (f *Firewall) simulate(stopIfCaught bool) {
	for pos, layer := range f.layers {
		// is there a scanner at the top of this layer?
		if layer != nil && layer.pos == 0 {
			// caught
			f.Severity += pos * layer.max
			f.caught = true
			if stopIfCaught {
				return
			}
		}
		// move scanners
		f.tick()
	}
}

// copy 2 firewalls.  assumes they have the same size of layers (ie, one has
// to have originally been a dup() of the other)
func copyFirewallState(dst, src *Firewall) {
	for i, srcLayer := range src.layers {
		if srcLayer != nil {
			*dst.layers[i] = *srcLayer
		}
	}
	dst.Severity = src.Severity
	dst.caught = src.caught
}

func (f *Firewall) dup() *Firewall {
	f2 := &Firewall{}
	f2.layers = make([]*layer, len(f.layers))
	for i := 0; i < len(f.layers); i++ {
		if f.layers[i] != nil {
			f2.layers[i] = &layer{}
		}
	}
	copyFirewallState(f2, f)
	return f2
}

func (f *Firewall) Run() *Firewall {
	f.simulate(false)
	return f
}

// SneakyWaitTime calculates the number of picoseconds we'd need to wait
// before sending a packet which won't get caught
func (f *Firewall) SneakyWaitTime() int {
	// array of 2 firewalls for swapping base
	fw := make([]*Firewall, 2)
	fw[0] = f.dup() // state after (wait) picoseconds
	fw[1] = f.dup() // scratch space to try simulations
	for wait := 0; ; wait++ {
		// try running after `wait` picoseconds
		fw[1].simulate(true)
		if !fw[1].caught {
			return wait
		}
		// next loop, try waiting an extra picosecond
		fw[0].tick()
		copyFirewallState(fw[1], fw[0])
	}
}
