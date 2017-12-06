package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	sum := 0
	
	for i := 1; i < len(args) ; i++ {
		temp, _ := strconv.Atoi(args[i])
		sum = sum + temp
	}
	fmt.Println("Sum of the values are : ", sum)
}
