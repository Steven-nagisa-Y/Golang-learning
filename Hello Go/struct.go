package main

import (
	"fmt"
)

type matrix struct {
	r int
	c int
}

func  main()  {
	var ma = matrix{1, 2}
	fmt.Println(ma.r)
}