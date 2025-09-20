package main

import (
	"fmt"
	"math"
)

func loop() {
	var sum int = 0 // or sim := 0
	//for loop
	for i := 0; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	//while loop
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func if_condition(x float64) string {
	fmt.Println("called with:", x) //test call to check recursion usage
	if x < 0 {
		return if_condition(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func if_var_init(x, n, lim float64) float64 {
	//initializing v var inside the if statement itself (a good way for test cases and err checks)
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(if_condition(2), if_condition(-2))
}
