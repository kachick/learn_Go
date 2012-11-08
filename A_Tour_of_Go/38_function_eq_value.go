package main

import (
	"fmt"
	"math"
)

// 関数内関数定義も可能

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4)) // 5
}