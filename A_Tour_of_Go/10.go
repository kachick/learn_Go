package main

import "fmt"

// YARDで示すと、こんなかんじ。・・・イイネ！
// @param [int] sum
// @return [int] x
// @return [int] y
func split(sum int) (x, y int) {
	x = sum * 4/9
	y = sum - x
	return // return x, y と同じ。名前代入されてた場合はこうなるらしい。
}

func split2(sum int) (x int) {
	x = sum * 4/9
	y := sum - x
	x = x - y
	return // return 最初にxしか名前つけてないと、途中でyを使っていようが、return x扱いになる。
	       // 紛らわしいんで、個人的にはreturnの中身省略したくない
}


func main() {
	fmt.Println(split(17))
	fmt.Println(split2(17))
}