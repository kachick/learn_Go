package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// RubyでいうとこのHashみたいなものか
// キーにも値にも型指定する
var m map[string]Vertex

func main() {
	// structはnew、mapはmakeで生成する
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, 74.39967,
	}
	fmt.Println(m["Bell Labs"]) // {40.68433 74.39967}
	fmt.Println(m["BellLabs"]) // {}
}