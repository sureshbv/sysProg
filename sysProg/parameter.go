package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	params := os.Args
	minusI := false
	
	for i := 1; i < len(params) ; i++ {
		if strings.Compare(params[i],"-i") == 0 {
			minusI = true
		}
	}

	if minusI {
		fmt.Println("-i is present")
		fmt.Print("Do you want to proceed: y/n ")
		var answer string
		fmt.Scanln(&answer)
		fmt.Println("You entered",answer)
	} else {
		fmt.Println("The -i parameter is not set")
	}	
}
