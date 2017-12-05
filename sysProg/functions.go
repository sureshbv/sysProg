package main

import "fmt"

func minMax(x, y int) (int, int) {
	if x > y {
		min := y
		max := x
		return min, max
	} else {
		min := x
		max := y
		return min, max
	}
}

func namedMinMax ( x , y int) (min , max int) {
		if x > y {
			min, max = y, x
		} else {
			min, max = x, y
		}
	return
}


func main () {
	y := 4
	
	square := func(s int) int {
		return s * s
	}

	fmt.Println("The square of ",y,"is ",square(y))

	square = func(s int) int {
		return s+s
	}
	fmt.Println("The 2nd square of ",y,"is ",square(y))

	fmt.Println(minMax(15,6))
	fmt.Println(namedMinMax(15,6))
	min, max := namedMinMax(12, -1)
	fmt.Println(min,max)
}
