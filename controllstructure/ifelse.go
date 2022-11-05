package main

import "fmt"

func printStudantSituation(grade float64) {
	if grade >= 7 {
		fmt.Println("Approved with grade: ", grade)
	} else {
		fmt.Println("Failed with grade: ", grade)
	}
}

func main() {
	printStudantSituation(7.5)
	printStudantSituation(5.1)
}
