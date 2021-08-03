package main

import (
	"fmt"
	// "math"
	// "math/rand"
)

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//[VARIABLE AND IMPORT]--------------------------------------------
	// fmt.Println("My fav number is: ", rand.Intn(10))
	// fmt.Printf("Now you have %g problem. \n", math.Sqrt(7))
	//-----------------------------------------------------------------

	// [FUNCTION]
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
