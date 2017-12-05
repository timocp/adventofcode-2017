package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

// CorruptionChecksum calculates the checksum of a spreadsheet
func CorruptionChecksum(input [][]int) (output int) {
	for _, row := range input {
		min := row[0]
		max := row[0]
		for _, v := range row {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
		}
		output += max - min
	}
	return
}

// EvenlyDivisibleChecksum is based on the 2 cells which evenly divide
func EvenlyDivisibleChecksum(input [][]int) (output int) {
	for _, row := range input {
		for i := range row {
			for j := 0; j < len(row); j++ {
				if row[j] < row[i] && row[i]%row[j] == 0 {
					output += row[i] / row[j]
					break
				}
			}
		}
	}
	return
}

// ReadSpreadsheet takes an io.Reader input (eg string or file) and retuns
// a 2d array of integers
func ReadSpreadsheet(input io.Reader) ([][]int, error) {
	output := [][]int{}
	s := bufio.NewScanner(input)
	for s.Scan() {
		output = append(output, []int{})
		for _, v := range strings.Fields(s.Text()) {
			i, err := strconv.Atoi(v)
			if err != nil {
				return output, err
			}
			output[len(output)-1] = append(output[len(output)-1], i)
		}
	}
	return output, s.Err()
}
