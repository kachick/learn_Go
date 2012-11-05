package main

import (
	"fmt"
	"io"
	"bufio"
	"os"
	"strings"
)

func gets(ior io.Reader, separator byte) (input string) {
	reader := bufio.NewReader(ior)
	input, _ = reader.ReadString(separator)
	return
}

func main() {
	fmt.Println("Input this mark: -_-")

	input := strings.TrimRight(gets(os.Stdin, '\n'), "\n")
	for input != `-_-` {
		input = gets(os.Stdin, '\n')
		fmt.Print(input)
		fmt.Println(`>_<`)
		fmt.Println(input != `-_-`)
	}
}
