package day10

type circle struct {
	list     []int
	position int
	skipSize int
}

func (c *circle) act(length int) {
	reverse(c.list, c.position, length)
	c.position = (c.position + length + c.skipSize) % len(c.list)
	c.skipSize++
}

func reverse(list []int, start int, length int) {
	// build up new values by iterating backwards
	tmp := make([]int, length)
	for i := 0; i < length; i++ {
		tmp[len(tmp)-1-i] = list[(start+i)%len(list)]
	}
	// now overwrite them forwards
	for i := 0; i < length; i++ {
		list[(start+i)%len(list)] = tmp[i]
	}
}

// Hash calculates the knot hash described in day 10 puzzle
func Hash(size int, lengths []int) int {
	c := &circle{make([]int, size), 0, 0}
	for i := range c.list {
		c.list[i] = i
	}
	for _, length := range lengths {
		c.act(length)
	}
	return c.list[0] * c.list[1]
}
