package main

import  (
	"fmt"
	"math"
)

const (
	Pi = 3.1428

	// Create a huge number by shifting a 1 bit left 100 places
	// In other words, the binary number that is 1 followed by
	// zeros.
	Big = 1 << 100
	// Shift if right again 99 places, so we end up
	// with 1 << 1, or 2
	Small = Big >> 99
)

func needInt(x int) int {
	return x * 10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x * x + y * y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v:= 42.0 // change me!
	fmt.Printf("Type: '%T' Value: %v\n", v, v)

	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)

	// A Tour of Go 16 of 17
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
	//fmt.Println(needInt(Big))
}
