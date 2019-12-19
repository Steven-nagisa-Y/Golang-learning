package main

import (
	"fmt"
)

// Define some variables
const (
	MAXI = 10
)

func main()  {
	var (
		i int
		p *int
	)
	for i=0; i<MAXI; i++ {
		p = &i
		fmt.Println(*p)
	}
}