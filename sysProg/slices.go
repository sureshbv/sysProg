package main

import (
	"fmt"
)

func printSlice(x []int) {

	for _, number := range x {
		fmt.Printf("%d ", number)
	}
	fmt.Println()
}

func change(x []int) {
	x[3] = -2
}


func main() {
	aSlice := []int{ 0, 1, -1, 4, 5, 6}
	fmt.Printf("Before change: ")
	printSlice(aSlice)
	change(aSlice)
	fmt.Printf("After change: ")
	printSlice(aSlice)
	fmt.Printf("Before. Cap: %d, length: %d \n", cap(aSlice), len(aSlice))
	aSlice = append(aSlice, -100)
	printSlice(aSlice)
	fmt.Printf("After. Cap: %d, length: %d \n", cap(aSlice), len(aSlice))
	anotherSlice := make([]int, 4)
	printSlice(anotherSlice)

}
