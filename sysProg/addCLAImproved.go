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
		temp, err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Println("Ignoring", args[i])
		} else {
			sum = sum + temp
		}
	} 
	fmt.Println("Sum of the values are : ", sum)
}
