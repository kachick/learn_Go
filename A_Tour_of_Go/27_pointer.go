package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	p := Vertex{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(q) // &{1000000000 2}
	fmt.Println(p) // {1000000000 2}
}