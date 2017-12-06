package main

import (
	"bufio"
	"io"
	"log"
	"strings"
)

// IsValidPassphrase returns true if the string contains no repeated words
func IsValidPassphrase(p string) bool {
	words := make(map[string]bool)
	for _, w := range strings.Fields(p) {
		if words[w] {
			return false
		}
		words[w] = true
	}
	return true
}

// CountValidPasshrases returns the number of valid passphrases found in input
func CountValidPasshrases(input io.Reader) (int, error) {
	count := 0
	s := bufio.NewScanner(input)
	for s.Scan() {
		if IsValidPassphrase(s.Text()) {
			count++
		}
	}
	return count, s.Err()
}

// MustCountValidPassphrases calls CountValidPasshrases with fatals on error
func MustCountValidPassphrases(input io.Reader) int {
	count, err := CountValidPasshrases(input)
	if err != nil {
		log.Fatal(err)
	}
	return count
}
