package main

import (
	"fmt"
)


const (
	// MAXI : what is this?
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