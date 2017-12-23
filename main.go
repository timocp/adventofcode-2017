package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/timocp/adventofcode/day1"
	"github.com/timocp/adventofcode/day10"
	"github.com/timocp/adventofcode/day11"
	"github.com/timocp/adventofcode/day12"
	"github.com/timocp/adventofcode/day13"
	"github.com/timocp/adventofcode/day14"
	"github.com/timocp/adventofcode/day15"
	"github.com/timocp/adventofcode/day16"
	"github.com/timocp/adventofcode/day17"
	"github.com/timocp/adventofcode/day18"
	"github.com/timocp/adventofcode/day19"
	"github.com/timocp/adventofcode/day2"
	"github.com/timocp/adventofcode/day20"
	"github.com/timocp/adventofcode/day21"
	"github.com/timocp/adventofcode/day23"
	"github.com/timocp/adventofcode/day3"
	"github.com/timocp/adventofcode/day4"
	"github.com/timocp/adventofcode/day5"
	"github.com/timocp/adventofcode/day6"
	"github.com/timocp/adventofcode/day7"
	"github.com/timocp/adventofcode/day8"
	"github.com/timocp/adventofcode/day9"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: adventofcode <puzzle> <inputfile>")
	}
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
		fmt.Println(day10.Knot(256, splitInts(os.Args[2])))
	case "10b":
		fmt.Println(day10.Hash(os.Args[2]))
	case "11a":
		fmt.Println(day11.ShortestDistance(mustReadFile(os.Args[2])))
	case "11b":
		fmt.Println(day11.MaxDistance(mustReadFile(os.Args[2])))
	case "12a":
		fmt.Println(day12.MustReadProgramList(mustOpen(os.Args[2])).CountConnected(0))
	case "12b":
		fmt.Println(day12.MustReadProgramList(mustOpen(os.Args[2])).CountGroups())
	case "13a":
		fmt.Println(day13.MustReadFirewall(mustOpen(os.Args[2])).Run().Severity)
	case "13b":
		fmt.Println(day13.MustReadFirewall(mustOpen(os.Args[2])).SneakyWaitTime())
	case "14a":
		fmt.Println(day14.SquaresUsed("jxqlasbh"))
	case "14b":
		fmt.Println(day14.RegionsPresent("jxqlasbh"))
	case "15a":
		fmt.Println(day15.Judge(516, 190, 40000000, false))
	case "15b":
		fmt.Println(day15.Judge(516, 190, 5000000, true))
	case "16a":
		fmt.Println(day16.Dance("abcdefghijklmnop", mustReadFile(os.Args[2])))
	case "16b":
		fmt.Println(day16.LongDance("abcdefghijklmnop", mustReadFile(os.Args[2]), 1000000000))
	case "17a":
		sl := day17.NewSpinLock(toInt(os.Args[2]))
		sl.Spins(2017)
		fmt.Println(sl.AtRel(1))
	case "17b":
		fmt.Println(day17.NewSpinLock(toInt(os.Args[2])).StopValue(50000000))
	case "18a":
		fmt.Println(day18.Part1(mustOpen(os.Args[2])))
	case "18b":
		fmt.Println(day18.Part2(mustReadFile(os.Args[2])))
	case "19a":
		t := day19.NewTubes(mustReadFile(os.Args[2]))
		t.Walk()
		fmt.Println(t.Letters)
	case "19b":
		t := day19.NewTubes(mustReadFile(os.Args[2]))
		t.Walk()
		fmt.Println(t.Steps)
	case "20a":
		s, err := day20.NewSwarm(mustReadFile(os.Args[2]), false)
		mustNot(err)
		s.Run(500)
		fmt.Println(s.Closest())
	case "20b":
		s, err := day20.NewSwarm(mustReadFile(os.Args[2]), true)
		mustNot(err)
		s.Run(500)
		fmt.Println(s.CountPresent())
	case "21a":
		fmt.Println(day21.Part1(mustOpen(os.Args[2])))
	case "23a":
		fmt.Println(day23.Part1(mustOpen(os.Args[2])))
	case "23b":
		fmt.Println(day23.Part2())
	default:
		fmt.Printf("Puzzle %s unimplemented\n", os.Args[1])
	}
}

// helper functions mostly to log and exit on errors

func mustNot(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func mustOpen(fn string) *os.File {
	f, err := os.Open(fn)
	mustNot(err)
	return f
}

func mustReadFile(fn string) string {
	bytes, err := ioutil.ReadFile(fn)
	mustNot(err)
	return string(bytes)
}

func mustReadTrampoline(fn string) *day5.Trampoline {
	tramp, err := day5.NewReadTrampoline(mustOpen(fn))
	mustNot(err)
	return tramp
}

func mustReadMemoryBank(fn string) *day6.MemoryBank {
	mem, err := day6.NewMemoryBankRead(mustOpen(fn))
	mustNot(err)
	return mem
}

func mustReadTower(fn string) *day7.Node {
	tower, err := day7.ReadTower(mustOpen(fn))
	mustNot(err)
	return tower
}

func mustCountValidPassphrases(input io.Reader, secure bool) int {
	count, err := day4.CountValidPasshrases(input, secure)
	mustNot(err)
	return count
}

func mustReadProgram(fn string) day8.Program {
	p, err := day8.ReadProgram(mustOpen(fn))
	mustNot(err)
	return p
}

// loadSS reads a 2d int array from a filename, calling log.Fatal on error
func loadSS(fn string) [][]int {
	f, err := os.Open(fn)
	mustNot(err)
	defer f.Close()
	ss, err := day2.ReadSpreadsheet(f)
	mustNot(err)
	return ss
}

func toInt(arg string) int {
	i, err := strconv.Atoi(arg)
	mustNot(err)
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
