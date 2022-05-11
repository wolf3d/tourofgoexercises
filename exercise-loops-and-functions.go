package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	var z = 1.0
	const threshold = 0.0012
	for (x - z*z) > threshold {
		//for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * x)
		//fmt.Printf("z %v\n", z)
		//fmt.Println(x - z*z)
		//fmt.Println(z*z, x)
		//fmt.Printf("math.Sqrt -> %v\n", math.Sqrt(x))
		//if (x - z*z) < threshold {
		//	break
		//}
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
