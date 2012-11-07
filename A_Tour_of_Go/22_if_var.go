package main

import (
	"fmt"
	"math"
)


func pow(x, n, lim float64) float64 {
	// if直後、条件の前に文をかける。ここで宣言した変数は、ifのブロックローカルになる
	if v := math.Pow(x, n); v < lim {
		return v
	}
	// return v // undefined: v
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 9
		pow(3, 3, 20), // 20
	)
}