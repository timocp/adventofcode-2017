package main

import "testing"

func TestCaptcha(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	} {
		r := Captcha(tt.in)
		if r != tt.out {
			t.Errorf("Captcha(%s) => %d, want %d", tt.in, r, tt.out)
		}
	}
}

func TestCaptchaHalfway(t *testing.T) {
	for _, tt := range []struct {
		in  string
		out int
	}{
		{"1212", 6},
		{"1221", 0},
		{"123425", 4},
		{"123123", 12},
		{"12131415", 4},
	} {
		r := CaptchaHalfway(tt.in)
		if r != tt.out {
			t.Errorf("Captcha2(%s) => %d, want %d", tt.in, r, tt.out)
		}
	}
}
