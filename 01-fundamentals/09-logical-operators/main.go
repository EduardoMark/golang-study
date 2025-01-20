package main

import "fmt"

func shopping(work1, work2 bool) (bool, bool, bool) {
	buyTv50 := work1 && work2
	buyTv32 := work1 != work2
	buyIceCream := work1 || work2

	return buyTv50, buyTv32, buyIceCream
}

func main() {
	tv50, tv32, iceCream := shopping(false, true)

	fmt.Printf("tv50: %t, tv32: %t, sorvete: %t, saudável: %t",
		tv50, tv32, iceCream, !iceCream)
}