package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 条件無しswitchは、elsif的。このへんもRubyと同じ
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}