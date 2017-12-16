package day16

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func move(in, m string) string {
	if strings.HasPrefix(m, "s") {
		n := toInt(m[1:len(m)])
		return in[len(in)-n:len(in)] + in[0:len(in)-n]
	} else if strings.HasPrefix(m, "x") {
		p := strings.Split(m[1:len(m)], "/")
		p0 := toInt(p[0])
		p1 := toInt(p[1])
		tmp := []byte(in)
		tmp[p0], tmp[p1] = tmp[p1], tmp[p0]
		return string(tmp)
	} else if strings.HasPrefix(m, "p") {
		p := strings.Split(m[1:len(m)], "/")
		p0 := strings.Index(in, p[0])
		p1 := strings.Index(in, p[1])
		tmp := []byte(in)
		tmp[p0], tmp[p1] = tmp[p1], tmp[p0]
		return string(tmp)
	}
	panic(fmt.Errorf("Invalid move: %s", m))
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func Dance(in, movesList string) string {
	return doDance(in, splitMoves(movesList))
}

func doDance(in string, moves []string) string {
	for _, m := range moves {
		in = move(in, m)
	}
	return in
}

// REALLY perform the dance `repeats` times.  Used to generate test data,
// far too slow to solve the real problem
/*
func LongDance(in, movesList string, repeats int) string {
	moves := splitMoves(movesList)
	for i := 0; i < repeats; i++ {
		in = doDance(in, moves)
		fmt.Printf("after %d dances, state=%s\n", i+1, in)
	}
	return in
}
*/

func LongDance(in, movesList string, repeats int) string {
	// dance until we see a repeated position
	seen := make(map[string]int)
	moves := splitMoves(movesList)
	count := 0
	for count < repeats {
		if _, ok := seen[in]; ok {
			for k, v := range seen {
				if v == repeats%count {
					return k
				}
			}
		}
		seen[in] = count
		in = doDance(in, moves)
		count++
	}
	// no repeats were found before reaching final pos
	return in
}

func splitMoves(movesList string) []string {
	moves := strings.Split(movesList, ",")
	for i := range moves {
		moves[i] = strings.TrimSpace(moves[i])
	}
	return moves
}
