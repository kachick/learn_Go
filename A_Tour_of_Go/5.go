package main

// import "fmt"
// import "math"
// の2行を、まとめて書く記法

import (
	"fmt"
	"math"
)

// math.Nextafterって？
// http://golang.org/src/pkg/math/nextafter.go

func main() {
	fmt.Printf("now you have %g prolems.",
		math.Nextafter(2, 3))
}