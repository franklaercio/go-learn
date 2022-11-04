package main

import "fmt"

func main() {
	i := 1

	// Go has no pointer arithmetic

	var p *int = nil

	p = &i

	*p++
	i++

	fmt.Println(p, *p, i, &i)
}
