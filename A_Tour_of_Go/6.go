// Exported names

package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(math.pi) #=> cannot refer to unexported name math.pi
	//                     #=> undefined: math.pi
	// Rubyでいうと、小文字はprivate、大文字はpublicになっているということか
	// これはたまたまでなく、常にそういう命名法則をとっているらしい。
	// Pythonの _var=private みたいなものかな
	fmt.Println(math.Pi)
}