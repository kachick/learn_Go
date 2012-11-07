package main

import "fmt"

func main() {
	sum := 0
	// for 直後に、()をつけてはいけない
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}