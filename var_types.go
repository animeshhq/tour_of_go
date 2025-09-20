/*
var types in go :
basic - number, string, bool etc
aggregate - array, struct etc
reference - pointers, slices, functions, channels, maps etc
interfaces
*/

package main

import (
	"fmt"
	"math/cmplx"
)

var python, js, golang bool // we need to use var when it comes to writing package level variables

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	i := 10 //this is same as writing var i int = 10, but since we are in a function level, we can use this ":=" as a syntatic sugar
	fmt.Printf("python: Type: %T Value: %v | js: Type: %T Value: %v | golang: Type: %T Value: %v\n", python, python, js, js, golang, golang)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
