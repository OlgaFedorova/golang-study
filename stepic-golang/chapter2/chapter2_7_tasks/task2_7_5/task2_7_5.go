package main

import (
	"fmt"
	"math"
)

var k, p, v float64

func main() {
	k = 1296
	p = 6
	v = 6
	fmt.Println(T())
}

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}
