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

func Dance(in, moves string) string {
	for _, m := range strings.Split(moves, ",") {
		in = move(in, strings.TrimSpace(m))
	}
	return in
}
