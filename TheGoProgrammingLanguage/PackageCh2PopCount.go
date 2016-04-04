package main

import (
	"gopl.io/ch2/popcount"
	"fmt"
)
func main()  {
	var x uint64 = 515
	fmt.Printf("%o  %o   %x\n", x, popcount.PopCount(x), x)
	a := byte(x>>(0*8))
	fmt.Println(a)
	a = byte(x>>(1*8))
	fmt.Println(a)
}

