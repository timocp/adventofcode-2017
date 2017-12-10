package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/timocp/adventofcode/day8"
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
	case "3b":
		fmt.Println(SpiralMemoryStressTest(toInt(os.Args[2])))
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
		mustReadProgram(os.Args[2]).Execute(&reg)
		fmt.Println(reg.LargestValue())
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

func mustReadTrampoline(fn string) *Trampoline {
	tramp, err := NewReadTrampoline(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return tramp
}

func mustReadMemoryBank(fn string) *MemoryBank {
	mem, err := NewMemoryBankRead(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return mem
}

func mustReadTower(fn string) *Node {
	tower, err := ReadTower(mustOpen(fn))
	if err != nil {
		log.Fatal(err)
	}
	return tower
}

func mustCountValidPassphrases(input io.Reader, secure bool) int {
	count, err := CountValidPasshrases(input, secure)
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
