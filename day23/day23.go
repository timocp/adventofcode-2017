package day23

import (
	"io"
	"log"

	"github.com/timocp/adventofcode/day18"
)

func Part1(input io.Reader) int {
	d := day18.NewDuet()
	err := d.Load(input)
	if err != nil {
		log.Fatal(err)
	}
	for d.Runable() {
		d.Exec(func(sound int) {})
	}
	return d.InvocationCount("mul")
}
