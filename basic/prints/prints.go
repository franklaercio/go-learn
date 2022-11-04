package main

import "fmt"

func main() {
	fmt.Print("Test 1")

	fmt.Println("Test 2")
	fmt.Println("Test 3")

	x := 3.141516

	xs := fmt.Sprint(x)

	fmt.Println("The values of x is " + xs)
	fmt.Println("The value of x is", x)

	fmt.Printf("The value of x is %f", x)
}
