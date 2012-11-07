package main

import "fmt"

// 定数はconstキーワードを使う
// character, string, boolean, (numeric values) でだけ使える
const Pi = 3.14

// 小文字も使える
const little_pi = 3

func main() {
	const World = "世界"
	fmt.Println("Hello, World")
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Happy", little_pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}