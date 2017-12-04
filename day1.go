package main

// Captcha based on each digit which matches the next in the string
func Captcha(input string) (output int) {
	p := input[len(input)-1]
	for i := 0; i < len(input); i++ {
		c := input[i]
		if p == c {
			output = output + int(p) - 48
		}
		p = c
	}
	return
}

// CaptchaHalfway compares characters halway around the string
func CaptchaHalfway(input string) (output int) {
	for i := 0; i < len(input); i++ {
		c := input[i]
		p := input[(i+len(input)/2)%len(input)]
		if p == c {
			output = output + int(p) - 48
		}
	}
	return
}
