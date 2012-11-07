package main

import "fmt"

// Numeric Constants

// constでも、まとめて書く記法が使える
const (
	Big = 1<<100
	Small = Big>>99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x*0.1
}

func main() {
	// fmt.Println(Small) // 2
	// fmt.Println(Big) //  error: constant 1267650600228229401496703205376 overflows int // <- これはエラーなのに

	fmt.Println(needInt(Small)) // 21
	fmt.Println(needFloat(Small)) // 0.2
	fmt.Println(needFloat(Big)) // 1.2676506002282295e+29 // <- これは平気ですよーという。(まだよく理解できていない)
}