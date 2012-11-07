package main

import "fmt"

func sum_all(min, max int) int {
	sum := 0
	for i := min; i < max; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println(sum_all(1, 10))
}