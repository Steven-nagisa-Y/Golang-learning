package main

import "fmt"

func main() {
	var i int
	for i=0; i<10; i++ {
		defer fmt.Printf("%d ",i)
	}
	fmt.Println("i is ", i)
}