package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1a":
		fmt.Println(Captcha(os.Args[2]))
	case "1b":
		fmt.Println(CaptchaHalfway(os.Args[2]))
	default:
		fmt.Printf("Puzzle %s unimplemented", os.Args[1])
	}
}
