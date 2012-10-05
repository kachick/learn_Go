package main

import "fmt"

func main() {
    loops := 1
    // while loop:
    for loops > 0 {
        fmt.Printf("\nNumber of loops?\n")
        fmt.Scanf("%d", &loops)
        // for loop
        for i := 0 ; i < loops ; i++ {
            fmt.Printf("%d", 1)
        }
    }

    // infinite
    for {
        // 明示的に終了! >_<
        break
    }
}
