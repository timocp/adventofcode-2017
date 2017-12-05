package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	switch os.Args[1] {
	case "1a":
		fmt.Println(InverseCaptcha(os.Args[2], 0))
	case "1b":
		fmt.Println(InverseCaptcha(os.Args[2], 1))
	case "2a":
		fmt.Println(CorruptionChecksum(loadSS(os.Args[2])))
	case "2b":
		fmt.Println(EvenlyDivisibleChecksum(loadSS(os.Args[2])))
	case "3a":
		fmt.Println(SpiralMemoryDistance(toInt(os.Args[2])))
	default:
		fmt.Printf("Puzzle %s unimplemented", os.Args[1])
	}
}

// loadSS reads a 2d int array from a filename, calling log.Fatal on error
func loadSS(fn string) [][]int {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ss, err := ReadSpreadsheet(f)
	if err != nil {
		log.Fatal(err)
	}
	return ss
}

func toInt(arg string) int {
	i, err := strconv.Atoi(arg)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
