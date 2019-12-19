package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	fmt.Println("Time ", time.Now())
	fmt.Println("number is:", rand.Intn(10))
}