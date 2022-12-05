package main

import (
	"fmt"
	"time"
)

func grade(g float64) string {
	var grade = int(g)

	switch grade {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1, 0:
		return "E"
	default:
		return "Invalid grade"
	}
}

func main() {
	fmt.Println(grade(10))
	fmt.Println(grade(7))
	fmt.Println(grade(5))
	fmt.Println(grade(2))
	fmt.Println(grade(-1))

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good Night!")
	}
}
