package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return (x-(math.Sqrt(x)-x)/(2*x))
}

func main() {
	fmt.Println(Sqrt(2))
}
