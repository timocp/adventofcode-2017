package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	switch os.Args[1] {
	case "1a":
		fmt.Println(InverseCaptcha(os.Args[2], 0))
	case "1b":
		fmt.Println(InverseCaptcha(os.Args[2], 1))
	case "2a":
		f, err := os.Open(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		ss, err := ReadSpreadsheet(f)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(CorruptionChecksum(ss))
	default:
		fmt.Printf("Puzzle %s unimplemented", os.Args[1])
	}
}
