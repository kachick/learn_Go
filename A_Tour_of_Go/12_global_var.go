package main

import "fmt"

// こんな感じで代入
var x, y, z int = 1, 2, 3
// 型省略すると、初期化時の型が使われる。
// var構文自体で型縛るわけじゃないので、こういう多重代入で型違いも通る
var c, python, java = true, false, "no!"

// 関数外では、:=が使えない
// c, python, java := true, false, "no!" //  non-declaration statement outside function body

func main() {
	fmt.Println(x, y, z, c, python, java)
}