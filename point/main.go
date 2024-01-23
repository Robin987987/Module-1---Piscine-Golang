package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	// fmt.Printf("x = %d, y = %d\n", points.x, points.y)
	s := "x = "
	s2 := ", y = "

	for _, character := range s {
		z01.PrintRune(character)
	}
	tempX := points.x
	tempY := points.y
	z01.PrintRune(rune(tempX/10) + 48)
	z01.PrintRune(rune(tempX%10) + 48)
	for _, character := range s2 {
		z01.PrintRune(character)
	}
	z01.PrintRune(rune(tempY/10) + 48)
	z01.PrintRune(rune(tempY%10) + 48)
	z01.PrintRune('\n')
}
