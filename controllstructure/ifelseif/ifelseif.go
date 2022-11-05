package main

import "fmt"

func studantSituation(grade float64) string {
	if grade >= 9 && grade <= 10 {
		return "A"
	} else if grade >= 8 && grade < 9 {
		return "B"
	} else if grade >= 5 && grade < 8 {
		return "C"
	} else if grade >= 3 && grade < 5 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	fmt.Println(studantSituation(10.0))
	fmt.Println(studantSituation(9.0))
	fmt.Println(studantSituation(8.0))
	fmt.Println(studantSituation(7.0))
	fmt.Println(studantSituation(6.0))
	fmt.Println(studantSituation(5.0))
	fmt.Println(studantSituation(4.0))
	fmt.Println(studantSituation(3.0))
	fmt.Println(studantSituation(2.0))
	fmt.Println(studantSituation(1.0))
}
