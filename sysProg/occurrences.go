package main

import (
	"fmt"
	"strings"
)


func main() {
	var s[3] string
	s[0] = "1 2 3"
	s[1] = "a b 2 d 3 f 1 h"
	s[2] = "3 b l c d e 1"

	counts := make(map[string]int)
	
	for i := 0; i < len(s); i++ {
		data := strings.Fields(s[i])
		
		for _,word := range data {
			_, ok := counts[word]
			if ok {
				counts[word] = counts[word] + 1
			} else {
				counts[word] = 1
			} 
		}
	}

	fmt.Println(counts)

	for key, _ := range counts {
		fmt.Printf("%s -> %d \n", key, counts[key])
	}
	
}

