package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(999))

	for i:=0; i < 10; i++ {
		bytes := rand.Intn(10)
		fmt.Print(bytes)
	}
}
