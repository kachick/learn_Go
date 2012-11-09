package main

import (
	"tour/wc"
)

// なんでintじゃなくmapを返す設計なのかようわからん。今度また考える

func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	wc.Test(WordCount)
}
