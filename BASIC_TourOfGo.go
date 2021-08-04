package main

import (
	"fmt"
	"math/cmplx"
	// "math"
	// "math/rand"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	myname string     = "ai"
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
	// fmt.Println(add(42, 13))
	// a, b := swap("hello", "world")
	// fmt.Println(a, b)

	//[BASIC TYPES]
	// fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v \n", z, z)
	// fmt.Printf("Type: %T Value: %v \n", z, z)
	// fmt.Printf("Type: %T Value: %v %q \n", myname, myname, myname)

	// [TYPE CONVERSION]
	i := 42.3
	f := int(i)
	fmt.Println(i)
	fmt.Println(f)

}
