package main

import (
	"fmt"
	"math"
)

// ニュートン法・・・挫折。今度再挑戦

func Sqrt(x float64) float64 {
	z := (x / 2) + 1
	last := 0.0
	for (z - last) >= 0.1 {
		fmt.Println(last, z)
		last = z
		z = z - (((z * z) - x) / 2 * z)
	}
	return z
}

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		fmt.Println(Sqrt(float64(i)))
		fmt.Println(math.Sqrt(float64(i)))
		fmt.Println("------------------------------------------------------------------")
	}
}