package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// 二つ返せるのはわかった。
	// でもそれって、Rubyの多重代入と何が違うんだ？
	// 例えば、つぎのような数違い代入がコンパイルエラーになる
	// c := swap("hello", "world") // multiple-value swap() in single-value context
	// d, e, f := swap("hello", "world") // assignment count mismatch: 3 = 2
}