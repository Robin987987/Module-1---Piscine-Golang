package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(arguments) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	textFile, err := ioutil.ReadFile(arguments[1])
	if err != nil {
		fmt.Println("Err")
	}
	fmt.Printf(string(textFile))
}
