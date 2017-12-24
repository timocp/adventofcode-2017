package day24

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
)

type port struct {
	a, b int
}

func parsePort(in string) port {
	pins := strings.Split(in, "/")
	return port{toInt(pins[0]), toInt(pins[1])}
}

type Bridge struct {
	ports *[]port
	used  []int
	end   int
}

func (b *Bridge) String() string {
	s := ""
	for i, pnbr := range b.used {
		if i > 0 {
			s += "--"
		}
		port := (*b.ports)[pnbr]
		s += fmt.Sprintf("%d/%d", port.a, port.b)
	}
	s += fmt.Sprintf(" {%d}", b.end)
	return s
}

func NewBridge(input io.Reader) *Bridge {
	b := &Bridge{}
	ports := []port{}
	s := bufio.NewScanner(input)
	for s.Scan() {
		ports = append(ports, parsePort(s.Text()))
	}
	if s.Err() != nil {
		log.Fatal(s.Err())
	}
	b.ports = &ports
	return b
}

func (b *Bridge) dup() *Bridge {
	new := &Bridge{b.ports, make([]int, len(b.used)), b.end}
	for i, p := range b.used {
		new.used[i] = p
	}
	return new
}

// the first call to this is an EMPTY bridge which we don't want to yield
func (b *Bridge) each(f func(*Bridge)) {
	for i, port := range *b.ports {
		if port.a == b.end || port.b == b.end {
			// would match, check it is not used
			used := false
			for _, pbr := range b.used {
				if pbr == i {
					used = true
					break
				}
			}
			if !used {
				// make a new bridge, yield it, then call recursively
				b2 := b.dup()
				b2.used = append(b2.used, i)
				if b.end == port.a {
					b2.end = port.b
				} else {
					b2.end = port.a
				}
				f(b2)
				b2.each(f)
			}
		}
	}
}

func (b *Bridge) strength() int {
	str := 0
	for _, pnbr := range b.used {
		port := (*b.ports)[pnbr]
		str += port.a + port.b
	}
	return str
}

func (b *Bridge) MaxStrength() int {
	maxStr := 0
	b.each(func(b2 *Bridge) {
		s2 := b2.strength()
		if s2 > maxStr {
			maxStr = s2
		}
	})
	return maxStr
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
