package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Sum ", a+b)
	fmt.Println("Sub ", a-b)
	fmt.Println("Div ", a/b)
	fmt.Println("Mul ", a*b)
	fmt.Println("Mod ", a%b)

	fmt.Println("And ", a&b) // 11 & 10 = 10
	fmt.Println("Or ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor ", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	fmt.Println("Greater ", math.Max(float64(c), float64(d)))
	fmt.Println("Smaller ", math.Min(float64(c), float64(d)))
	fmt.Println("Exponent ", math.Pow(c, d))
}
