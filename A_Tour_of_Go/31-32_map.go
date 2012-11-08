package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// いきなり生成
// MapもVertexも{}で使うのが若干分かりづらい
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408, //　最後のカンマ外すとエラー。改行しなけりゃ無くてもOKっぽい
	},
}

var m2 = map[string]Vertex{
	"Bell Labs": { // 型名省略可
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}


func main() {
	fmt.Println(m) // map[Google:{37.42202 -122.08408} Bell Labs:{40.68433 -74.39967}]
	fmt.Println(m2) // map[Google:{37.42202 -122.08408} Bell Labs:{40.68433 -74.39967}]
}