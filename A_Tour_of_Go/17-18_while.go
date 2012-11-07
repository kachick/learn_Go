package main

import "fmt"

func main() {
	sum := 1
	// for の前後処理は省略し、継続条件のみにすることが出来る
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)

	sum2 := 1
	// セミコロンも省略できる。というか、要するにwhile文である
	// ちなみに、省略する時は二つとも省略すること
	// for ; sum < 1000 や
	// for sum < 1000; はエラー
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
}