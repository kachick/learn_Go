package main

import (
 	"fmt"
 	"math"
)

func main() {
	fmt.Println("Happy", math.Pi, "Day")
}

// importしたmathが、そのままmath.Piとしての名前空間になっている。これは仕様というより、ルール。（マナーよりは厳しい？）