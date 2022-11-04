package main

import (
	"fmt"
	"math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 //Base64

	// Assignment value in Go
	area := PI * math.Pow(raio, 2)

	fmt.Println("The circumference area is: ", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	var e, f bool = true, false
	fmt.Println(e, f)

	// Assignment value in Go
	g, h, i := 2, false, "test!"
	fmt.Println(g, h, i)
}
