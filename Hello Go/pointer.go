package main

import (
	"fmt"
)
const (
	maxi = 10
)

func main()  {
	var (
		i int
		p *int
	)
	for i=0; i<maxi; i++ {
		p = &i
		fmt.Println(*p)
	}
}