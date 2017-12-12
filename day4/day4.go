package day4

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

// IsValidPassphrase returns true if the string contains no repeated words
// if secure is true, no word can be an anagram of a previous word
func IsValidPassphrase(p string, secure bool) bool {
	words := make(map[string]bool)
	for _, w := range strings.Fields(p) {
		if secure {
			w = sortString(w)
		}
		if words[w] {
			return false
		}
		words[w] = true
	}
	return true
}

// CountValidPasshrases returns the number of valid passphrases found in input
func CountValidPasshrases(input io.Reader, secure bool) (int, error) {
	count := 0
	s := bufio.NewScanner(input)
	for s.Scan() {
		if IsValidPassphrase(s.Text(), secure) {
			count++
		}
	}
	return count, s.Err()
}

func sortString(in string) string {
	s := strings.Split(in, "")
	sort.Strings(s)
	return strings.Join(s, "")
}
