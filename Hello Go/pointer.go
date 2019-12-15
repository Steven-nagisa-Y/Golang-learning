package main

import "fmt"
var (
	NAME = "JH"
	AGE = 19
)

func main() {
	p := &NAME
	fmt.Println(*p)
}