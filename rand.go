package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main2() {
	rand.Seed(time.Now().Unix())
	fmt.Println("My favorite number is", rand.Intn(10))
}
