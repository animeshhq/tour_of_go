//a different way of returning values in go (not recommened to use this method)

package main

import (
	"fmt"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
