package main

import "fmt"

// 関数外では、「すべての宣言は、キーワードから始まる」

func main() {
	var x, y, z int = 1, 2, 3
	// 関数内でのみ、 := での初期化が可能
	c, python, java := true, false, "no!"

	fmt.Println(x, y, z, c, python, java)
}