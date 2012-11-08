package main

import "fmt"

// sliceは、配列そのものではなく、"「配列っぽく扱える」「配列や、その部分配列への参照」"っぽい。とりあえずそう理解しておく

func main() {
	p := []int{2, 3, 5, 7, 11, 13} // intだけのslice宣言
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ { // len で長さを返す
		fmt.Println("p[%d] == %d\n",
			i, p[i])
	}
}