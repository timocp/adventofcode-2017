package day14

import (
	"fmt"
	"math/bits"

	"github.com/timocp/adventofcode/day10"
)

func SquaresUsed(input string) int {
	used := 0
	for row := 0; row < 128; row++ {
		for _, b := range day10.HashValue(fmt.Sprintf("%s-%d", input, row)) {
			used += bits.OnesCount8(uint8(b))
		}
	}
	return used
}
