package main

import "fmt"

func main() {
	// Same date in array
	var grades [3]float64
	grades[0], grades[1], grades[2] = 6.0, 10.0, 8.5

	fmt.Printf("Grades %.2f\n", grades)

	total := 0.0

	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}

	average := total / float64(len(grades))
	fmt.Printf("Average %.2f\n", average)
}
