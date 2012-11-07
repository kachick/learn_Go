package main

import "fmt"

// 宣言だけして値渡していない変数は、int: 0,bool: false で決め打ちされる
var x, y, z int
var c, python, java bool

func main() {
	fmt.Println(x, y, z, c, python, java)
}