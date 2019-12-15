package main

import (
	"fmt"
)

const (
	POOL = 100
)

func main() {
	num := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(num[:])
	var str [POOL]string
	fmt.Println(len(str))
}