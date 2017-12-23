package day23

import (
	"io"
	"log"
	"math/big"

	"github.com/timocp/adventofcode/day18"
)

func Part1(input io.Reader) int {
	d := day18.NewDuet()
	mustLoad(d, input)
	for d.Runable() {
		d.Exec(func(sound int) {})
	}
	return d.InvocationCount("mul")
}

/*
	a	constant 1
	b	starts at 108100
	c	constant 125100
	d	middle loop variable
	e	inner loop variable
	f	boolean (0 if any d divides b during middle loop)
	g	temporary register for comparisons
	h	counts the number of times the middle loop exits with f 0.
*/
func Part2() int {
	// the program is calculating the count of numbers between 108100..125100
	// (step 17) which are NOT primes.
	h := 0
	for b := int64(108100); b <= 125100; b += 17 {
		if !big.NewInt(b).ProbablyPrime(1) {
			h++
		}
	}
	return h
}

func mustLoad(d *day18.Duet, input io.Reader) {
	err := d.Load(input)
	if err != nil {
		log.Fatal(err)
	}
}
