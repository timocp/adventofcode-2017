package day10

import "fmt"

type circle struct {
	list     []byte
	position int
	skipSize int
}

func (c *circle) act(length int) {
	reverse(c.list, c.position, length)
	c.position = (c.position + length + c.skipSize) % len(c.list)
	c.skipSize++
}

func reverse(list []byte, start int, length int) {
	// build up new values by iterating backwards
	tmp := make([]byte, length)
	for i := 0; i < length; i++ {
		tmp[len(tmp)-1-i] = list[(start+i)%len(list)]
	}
	// now overwrite them forwards
	for i := 0; i < length; i++ {
		list[(start+i)%len(list)] = tmp[i]
	}
}

// Knot calculates one round of a knot hash described in day 10 puzzle
func Knot(size int, lengths []int) int {
	c := newCircle(size)
	c.round(lengths)
	return int(c.list[0]) * int(c.list[1])
}

func (c *circle) round(lengths []int) {
	for _, length := range lengths {
		c.act(length)
	}
}

func newCircle(size int) *circle {
	c := circle{make([]byte, size), 0, 0}
	for i := range c.list {
		c.list[i] = byte(i)
	}
	return &c
}

func Hash(input string) string {
	hash := HashValue(input)
	output := ""
	for _, v := range hash {
		output += fmt.Sprintf("%02x", v)
	}
	return output
}

func HashValue(input string) []byte {
	c := newCircle(256)
	lengths := inputToLengths(input)
	for i := 0; i < 64; i++ {
		c.round(lengths)
	}
	return denseHash(c.list)
}

// inputToLengths converts a string to bytes, then returns a list of them
// converted to ints.  the standard suffixes are appending to the slice.
func inputToLengths(input string) []int {
	// convert to ascii bytes first
	bytes := []byte(input)
	lengths := make([]int, len(bytes)+5)
	for i, b := range bytes {
		lengths[i] = int(b)
	}
	copy(lengths[len(bytes):], []int{17, 31, 73, 47, 23})
	return lengths
}

func denseHash(list []byte) []byte {
	d := make([]byte, 16)
	for i := 0; i < 16; i++ {
		d[i] = list[16*i]
		for j := 16*i + 1; j < 16*i+16; j++ {
			d[i] = d[i] ^ list[j]
		}
	}
	return d
}
