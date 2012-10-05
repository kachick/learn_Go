package main

const (
	Red	            = (1<<iota)
	Green           = (1<<iota)
	Blue, ColorMask = (1<<iota), (1<<(iota+1))-1
)

func main () {
}
