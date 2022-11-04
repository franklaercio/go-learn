package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(1, 3, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// No signal uint8 uint16 uint32 uint64

	var b byte = 255
	fmt.Println("O byte é", reflect.TypeOf(b)) // uint8

	var x float32 = 49.99
	fmt.Println("O tipo de x é ", reflect.TypeOf(x))                 // float32
	fmt.Println("O tipo de literal 49.99 é ", reflect.TypeOf(49.99)) // float64

	s1 := "Hello"
	fmt.Println("O tamanho da string é", len(s1)) //5
}
