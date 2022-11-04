package main

import "fmt"

func buy(work1, work2 bool) (bool, bool, bool) {
	buyTv50 := work1 && work2
	buyTv32 := work1 != work2
	buyIcecream := work1 || work2

	return buyTv50, buyTv32, buyIcecream
}

func main() {
	tv50, tv32, icecream := buy(true, true)
	fmt.Println("TV 50: %t, TV 32: %t, Icecream: %t", tv50, tv32, icecream)
}
