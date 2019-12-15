package main

import (
	"fmt"
)

type Matrix struct {
	r int
	c int
}

func  main()  {
	var ma = Matrix{1, 2}
	fmt.Println(ma.r)
}