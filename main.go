package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/timocp/adventofcode/day1"
	"github.com/timocp/adventofcode/day10"
	"github.com/timocp/adventofcode/day2"
	"github.com/timocp/adventofcode/day3"
	"github.com/timocp/adventofcode/day4"
	"github.com/timocp/adventofcode/day5"
	"github.com/timocp/adventofcode/day6"
	"github.com/timocp/adventofcode/day7"
	"github.com/timocp/adventofcode/day8"
	"github.com/timocp/adventofcode/day9"
)

func main() {
	switch os.Args[1] {
	case "1a":
		fmt.Println(day1.InverseCaptcha(os.Args[2], 0))
	case "1b":
		fmt.Println(day1.InverseCaptcha(os.Args[2], 1))
	case "2a":
		fmt.Println(day2.CorruptionChecksum(loadSS(os.Args[2])))
	case "2b":
		fmt.Println(day2.EvenlyDivisibleChecksum(loadSS(os.Args[2])))
	case "3a":
		fmt.Println(day3.SpiralMemoryDistance(toInt(os.Args[2])))
	case "3b":
		fmt.Println(day3.SpiralMemoryStressTest(toInt(os.Args[2])))
	case "4a":
		fmt.Println(mustCountValidPassphrases(mustOpen(os.Args[2]), false))
	case "4b":
		fmt.Println(mustCountValidPassphrases(mustOpen(os.Args[2]), true))
	case "5a":
		fmt.Println(mustReadTrampoline(os.Args[2]).StepsToExit(false))
	case "5b":
		fmt.Println(mustReadTrampoline(os.Args[2]).StepsToExit(true))
	case "6a":
		fmt.Println(mustReadMemoryBank(os.Args[2]).RedistributeUntilRepeat().Cycle)
	case "6b":
		fmt.Println(mustReadMemoryBank(os.Args[2]).RedistributeUntilRepeat().Repeat)
	case "7a":
		fmt.Println(mustReadTower(os.Args[2]).Name)
	case "7b":
		fmt.Println(mustReadTower(os.Args[2]).WrongWeightShouldBe())
	case "8a":
		reg := day8.NewRegisters()
		_ = mustReadProgram(os.Args[2]).Execute(&reg)
		fmt.Println(reg.LargestValue())
	case "8b":
		reg := day8.NewRegisters()
		fmt.Println(mustReadProgram(os.Args[2]).Execute(&reg))
	case "9a":
		fmt.Println(day9.MustProcess(mustOpen(os.Args[2])).Score)
	case "9b":
		fmt.Println(day9.MustProcess(mustOpen(os.Args[2])).GarbageCount)
	case "10a":
		fmt.Println(day10.Hash(256, splitInts(os.Args[2])))
	default:
		fmt.Printf("Puzzle %s unimplemented\n", os.Args[1])
	}
}

// helper functions mostly to log and exit on errors

func mustOpen(fn string) *os.File {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func mustReadTrampoline(fn string) *day5.Trampoline {
	tramp, err := day5.NewReadTrampoline(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return tramp
}

func mustReadMemoryBank(fn string) *day6.MemoryBank {
	mem, err := day6.NewMemoryBankRead(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return mem
}

func mustReadTower(fn string) *day7.Node {
	tower, err := day7.ReadTower(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return tower
}

func mustCountValidPassphrases(input io.Reader, secure bool) int {
	count, err := day4.CountValidPasshrases(input, secure)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func mustReadProgram(fn string) day8.Program {
	p, err := day8.ReadProgram(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return p
}

// loadSS reads a 2d int array from a filename, calling log.Fatal on error
func loadSS(fn string) [][]int {
	f, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	ss, err := day2.ReadSpreadsheet(f)
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

func splitInts(arg string) []int {
	words := strings.Split(arg, ",")
	ints := make([]int, len(words))
	for i, w := range words {
		ints[i] = toInt(w)
	}
	return ints
}
