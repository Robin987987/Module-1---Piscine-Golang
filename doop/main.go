package main

import (
	"os"
)

func main() {
	arguments := os.Args
	// Check if inputs are valid
	if len(arguments) < 4 {
		return
	}
	digits := "-0123456789"
	operators := "+-/*%"
	isD := false
	a := 0
	b := 0
	minus := false
	minus2 := false
	for _, digit := range arguments[1] {
		if digit != '-' {
			a *= 10
			a += int(digit) - 48
		} else {
			minus = true
		}
		isD = false
		for _, digit2 := range digits {
			if digit == digit2 {
				isD = true
			}
		}
		if !isD {
			return
		}
	}

	for _, digit := range arguments[3] {
		if digit != '-' {
			b *= 10
			b += int(digit) - 48
		} else {
			minus2 = true
		}
		isD = false
		for _, digit2 := range digits {
			if digit == digit2 {
				isD = true
			}
		}
		if !isD {
			return
		}
	}
	isO := false
	for _, operator2 := range operators {
		if rune(arguments[2][0]) == operator2 {
			isO = true
		}
	}
	if !isO {
		return
	}

	if !checkD(rune((a % 10) + 48)) {
		return
	}
	if !checkD(rune((b % 10) + 48)) {
		return
	}
	if minus {
		a *= -1
	}
	if minus2 {
		b *= -1
	}
	// Start calculating
	c := 0
	operator := arguments[2][0]
	if operator == '+' {
		c = a + b
		PrintDigit(c)
	}
	if operator == '-' {
		c = a - b
		PrintDigit(c)
	}
	if operator == '/' {
		if b == 0 {
			os.Stderr.WriteString("No division by 0")
			os.Stderr.WriteString("\n")
			return
		}
		c = a / b
		PrintDigit(c)
	}
	if operator == '%' {
		if b == 0 {
			os.Stderr.WriteString("No modulo by 0")
			os.Stderr.WriteString("\n")
			return
		}
		c = a % b
		PrintDigit(c)
	}
	if operator == '*' {
		if (a > (9223372036854775808/2) && b > 1) || (b > (9223372036854775808/2) && a > 1) {
			return
		}
		c = a * b
		PrintDigit(c)
	}
}

func PrintDigit(digit int) {
	digitSlice := ""
	minus := false
	if digit < 0 {
		digit *= -1
		minus = true
	}
	if digit == 0 {
		os.Stderr.WriteString("0")
		os.Stderr.WriteString("\n")
		return
	}
	for digit != 0 {
		digitSlice = string((digit%10)+48) + digitSlice
		digit /= 10
	}
	if !checkDs(digitSlice) {
		return
	}
	if minus {
		os.Stderr.WriteString("-")
	}
	os.Stderr.WriteString(digitSlice)
	os.Stderr.WriteString("\n")
}

func checkD(digit rune) bool {
	digits := "0123456789"
	isD := false
	for _, digit2 := range digits {
		if digit == digit2 {
			isD = true
		}
	}
	return isD
}

func checkDs(digitB string) bool {
	digits := "0123456789"
	isD := false
	for _, digit := range digitB {
		isD = false
		for _, digit2 := range digits {
			if digit == digit2 {
				isD = true
			}
		}
		if isD == false {
			return false
		}
	}
	return true
}
