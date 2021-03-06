package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println(sum)
	}
	fmt.Println(sum)

	var loops int = 1
	for sum < 1000 {
		sum += sum
		loops += 1
	}
	fmt.Printf("loops: %v sum: %v\n", loops, sum)

	fmt.Println(sqrt(2), sqrt(-4))

	// Flow control 6/14
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}
