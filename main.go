package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1a":
		fmt.Println(InverseCaptcha(os.Args[2], 0))
	case "1b":
		fmt.Println(InverseCaptcha(os.Args[2], 1))
	default:
		fmt.Printf("Puzzle %s unimplemented", os.Args[1])
	}
}
