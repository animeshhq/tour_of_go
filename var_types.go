package main

import (
	"fmt"
	"math/cmplx"
)

var python, js, golang bool

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	i := 10
	fmt.Printf("python: Type: %T Value: %v | js: Type: %T Value: %v | golang: Type: %T Value: %v\n", python, python, js, js, golang, golang)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
