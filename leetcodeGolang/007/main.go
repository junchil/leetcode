package main

import (
	"fmt"
)

func reverse(x int) int {
	sign := 1
	if x < 0 {
		sign = -1
		x = -1 * x
	}

	newint := 0

	for x > 0 {
		temp := x % 10
		newint = newint*10 + temp
		x = x / 10
	}

	newint = sign * newint

	return newint

}

func main() {
	x := 59987687
	y := reverse(x)
	fmt.Println(y)
}
