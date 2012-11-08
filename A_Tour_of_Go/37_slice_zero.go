package main

import "fmt"

// 「sliceのゼロの値はnilです。」とあるけど、どういうこっちゃ。emptyなsliceはnilということ？

func main() {
	var z []int
	fmt.Println(z, len(z), cap(z)) // [] 0 0
	if z == nil {
		fmt.Println("nil!")
	} // nil!
}
