package day1

// InverseCaptcha calcualtes the day1 capcha from an input string
// mode = 0 - each char is compared to the following char
// mode = 1 - each char is compared to one halfway around the input
func InverseCaptcha(input string, mode int) (output int) {
	offset := 1
	if mode == 1 {
		offset = len(input) / 2
	}
	for i := 0; i < len(input); i++ {
		c := input[i]
		p := input[(i+offset)%len(input)]
		if p == c {
			output += int(p) - 48
		}
	}
	return
}
