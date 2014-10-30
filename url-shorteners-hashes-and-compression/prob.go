package main

import (
	"fmt"
	"math"
)

func main() {
	n := math.Exp2(16)
	k := 1000
	x := float64((-1*k)*(k-1)) / float64(2*n)
	fmt.Printf("%v", 1-math.Exp(x))

}
