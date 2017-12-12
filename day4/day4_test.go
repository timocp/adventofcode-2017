package day4

import (
	"bytes"
	"testing"
)

func TestIsValidPassphrase(t *testing.T) {
	for _, tt := range []struct {
		in     string
		secure bool
		out    bool
	}{
		{"aa bb cc dd ee", false, true},
		{"aa bb cc dd aa", false, false},
		{"aa bb cc dd aaa", false, true},
		// part 2
		{"abcde fghij", true, true},
		{"abcde xyz ecdab", true, false},
		{"a ab abc abd abf abj", true, true},
		{"iiii oiii ooii oooi oooo", true, true},
		{"oiii ioii iioi iiio", true, false},
	} {
		r := IsValidPassphrase(tt.in, tt.secure)
		if r != tt.out {
			t.Errorf("IsValidPassphrase(%s, %t) => %t, want %t\n", tt.in, tt.secure, r, tt.out)
		}
	}
}

func TestCountValidPassphrases(t *testing.T) {
	for _, tt := range []struct {
		in     string
		secure bool
		out    int
	}{
		{"aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa\n", false, 2},
		{"abcde fghij\nabcde xyz ecdab\na ab abc abd abf abj\niiii oiii ooii oooi oooo\noiii ioii iioi iiio\n", true, 3},
	} {
		r, err := CountValidPasshrases(bytes.NewBufferString(tt.in), tt.secure)
		if err != nil {
			t.Errorf("CountValidPasshrases returned error %v, expected nil", err)
		} else if r != tt.out {
			t.Errorf("CountValidPasshrases(%s, %t) => %d, want %d", tt.in, tt.secure, r, tt.out)
		}
	}
}
