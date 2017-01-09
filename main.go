package main

import (
	"fmt"
	"math/rand"
	"time"
	"math"
	"math/cmplx"
)

func add(x, y int) int{
	return x + y
}

func swap(x, y string) (string, string){
	return y, x
}

func split(sum int) (x, y int){
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe	bool	 	= false
	MaxInt	uint64		= 1<<64 - 1
	z		complex128	= cmplx.Sqrt(-5 + 12i)
)

func main() {
	var c, python, java = true, false, "no?"
	var j, k  = 1, 2
	m := 3

	rand.Seed(time.Now().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(999))

	for i:=0; i < 10; i++ {
		bytes := rand.Intn(10)
		fmt.Print(bytes)
	}

	fmt.Println("\nThe time is", time.Now())
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(12, 5))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	fmt.Println(split(17))

	fmt.Println(j, k, m, c, python, java)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
