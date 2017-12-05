package main

import "math"

// SpiralMemoryDistance measures the manhatten distance from cell i to the
// centre of the memory grid
func SpiralMemoryDistance(i int) int {
	if i == 0 {
		return 0
	}
	n := int((math.Sqrt(float64(i-1)) + 1) / 2)
	//move := 0
	east := square((n*2)-1) + n

	// work out minimum move in this square to one of the cardinals; can't be
	// more than the square number
	move := n
	for cardinal := 0; cardinal < 4; cardinal++ {
		directPath := east + cardinal*n*2
		thisMove := iAbs(directPath - i)
		if thisMove < move {
			move = thisMove
		}
	}
	return n + move
}

func square(i int) int {
	return i * i
}

func iAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

/*
work out which square away from the middle we're in

sq	range
0	1..1
1	2..9
2	10..25
3	26..49
4	50..81
n	(((n*2)-1)^2)+1 .. ((n*2)+1)^2

=> cell x is in square (sqrt(n-1) + 1) / 2 (rounding down)

In square n (n>0), the most direct lines are:

sq	east			north			west			south
1	2				4				6				8
2	11				15				19				23
3	28				34				40				46
n	((n*2)-1)**2+n	east + n*2		north + n*2		north + n*2

If on these, the answer is the square number
If not, the minimum difference to one of the cardinals must be moved first
*/
