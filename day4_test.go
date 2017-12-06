package main

import (
	"bytes"
	"testing"
)

func TestIsValidPassphrase(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out bool
	}{
		{"aa bb cc dd ee", true},
		{"aa bb cc dd aa", false},
		{"aa bb cc dd aaa", true},
	} {
		r := IsValidPassphrase(tt.in)
		if r != tt.out {
			t.Errorf("IsValidPassphrase(%s) => %t, want %t\n", tt.in, r, tt.out)
		}
	}
}

func TestCountValidPassphrases(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{
		{"aa bb cc dd ee\naa bb cc dd aa\naa bb cc dd aaa\n", 2},
	} {
		r, err := CountValidPasshrases(bytes.NewBufferString(tt.in))
		if err != nil {
			t.Errorf("CountValidPasshrases returned error %v, expected nil", err)
		} else if r != tt.out {
			t.Errorf("CountValidPasshrases(%s) => %d, want %d", tt.in, r, tt.out)
		}
	}
}
