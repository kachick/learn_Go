package main

import "fmt"

type Vertex struct {
	X int // field
	Y int // field
}

func main() {
	fmt.Println(Vertex{1, 2}) // {1, 2}
	vertex := Vertex{1, 2}
	fmt.Println(vertex) // {1, 2}
	// vertex.Y = "8" // cannot use "8" (type string) as type int in assignment
	vertex.Y = 8
	fmt.Println(vertex) // {1, 8}
}