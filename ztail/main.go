package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args
	tailNbr := 0
	isTail := false
	tailLock := false
	exitStatus := false
	isFirst := true
	if len(arguments) < 4 {
		return
	}
	for _, s := range arguments[1:] {

		if !tailLock {
			isTail = true
			for _, character := range s {
				if 47 < character && character < 58 {
					tailNbr *= 10
					tailNbr += int(character) - 48
				} else {
					isTail = false
				}
			}
			if isTail {
				tailLock = true
			} else {
				tailNbr = 0
			}
		}
		if isTail {
			isTail = false
			continue
		}
		if s == "-c" {
			continue
		} else {
			content, err := os.ReadFile(s)
			if err != nil {
				exitStatus = true
				fmt.Printf("open %s: no such file or directory", s)
				fmt.Printf("\n")
				isFirst = false
				continue
			} else {
				if !isFirst {
					fmt.Printf("\n")
				}
				fmt.Printf("==> %s <==", s)
				fmt.Printf("\n")
				if len(string(content)) <= tailNbr {
					fmt.Printf(string(content))
				} else {
					fmt.Printf(string(content)[len(string(content))-tailNbr:])
				}
				isFirst = false
			}
		}
	}
	if exitStatus {
		os.Exit(1)
	}
}
