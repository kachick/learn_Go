package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1<<uint(i) // そもそもこの演算が良く分かってないなー。まぁ、map!的なあれなんだろうけど
	}
	for _, value := range pow { // _を使って捨てないと、「その後利用していませんよ」なコンパイルエラーになる。
		fmt.Printf("%d\n", value)
	}
}