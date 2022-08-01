package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("hello, world\n")
	var mandee string = "My real name is Amanda"
	fmt.Println(mandee)
	//I dont know why we need to use rand.Seed to get random numbers each time we use rand.Intn but
	// I would like to look into that more
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
