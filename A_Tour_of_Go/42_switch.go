package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	// 基本的にフォールスルーしないのは、Rubyと同じ
	switch os := runtime.GOOS; os { // os変数のスコープは、このcaseブロックローカル
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
	// fmt.Println(os) // ここでのos変数はundefined
}