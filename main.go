package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(999))

	for i:=0; i < 10; i++ {
		bytes := rand.Intn(10)
		fmt.Print(bytes)
	}

	fmt.Println("\nThe time is", time.Now())
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
}
