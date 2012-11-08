package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)
	fmt.Println("p[1:4] ==", p[1:4]) // {3, 5, 7}

	// missing low index implies 0
	fmt.Println("p[:3] ==", p[:3]) // {2, 3, 5}

	// missing high index implies len(s)
	fmt.Println("p[4:] ==", p[4:]) // {11, 13}
}